package export

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	kumactl_cmd "github.com/kumahq/kuma/app/kumactl/pkg/cmd"
	"github.com/kumahq/kuma/app/kumactl/pkg/output"
	"github.com/kumahq/kuma/app/kumactl/pkg/output/printers"
	"github.com/kumahq/kuma/app/kumactl/pkg/output/table"
	"github.com/kumahq/kuma/app/kumactl/pkg/output/yaml"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_system "github.com/kumahq/kuma/pkg/core/resources/apis/system"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
)

type exportContext struct {
	*kumactl_cmd.RootContext

	args struct {
		profile string
		format  string
	}
}

const (
	profileFederation             = "federation"
	profileFederationWithPolicies = "federation-with-policies"
	profileAll                    = "all"

	formatUniversal  = "universal"
	formatKubernetes = "kubernetes"
)

var staticProfiles = map[string][]model.ResourceType{
	profileFederation: {
		core_mesh.MeshType,
		core_system.GlobalSecretType,
		core_system.SecretType,
	},
}

func NewExportCmd(pctx *kumactl_cmd.RootContext) *cobra.Command {
	ctx := &exportContext{RootContext: pctx}
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export Kuma resources",
		Long:  `Export Kuma resources.`,
		Example: `
Export Kuma resources
$ kumactl export --profile federation --format universal > policies.yaml
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := pctx.CheckServerVersionCompatibility(); err != nil {
				cmd.PrintErrln(err)
			}

			if ctx.args.profile != profileAll && ctx.args.profile != profileFederation && ctx.args.profile != profileFederationWithPolicies {
				return errors.New("invalid profile")
			}

			resTypes, err := resourcesTypesToDump(cmd.Context(), ctx)
			if err != nil {
				return err
			}

			if ctx.args.format != formatUniversal && ctx.args.format != formatKubernetes {
				return errors.New("invalid format")
			}

			rs, err := pctx.CurrentResourceStore()
			if err != nil {
				return err
			}

			meshes := &core_mesh.MeshResourceList{}
			if err := rs.List(cmd.Context(), meshes); err != nil {
				return errors.Wrap(err, "could not list meshes")
			}

			var resources []model.Resource
			for _, resType := range resTypes {
				resDesc, err := pctx.Runtime.Registry.DescriptorFor(resType)
				if err != nil {
					return err
				}
				if resDesc.Scope == model.ScopeGlobal {
					list := resDesc.NewList()
					if err := rs.List(cmd.Context(), list); err != nil {
						return errors.Wrapf(err, "could not list %q", resType)
					}
					resources = append(resources, list.GetItems()...)
				} else {
					for _, mesh := range meshes.Items {
						list := resDesc.NewList()
						if err := rs.List(cmd.Context(), list, store.ListByMesh(mesh.GetMeta().GetName())); err != nil {
							return errors.Wrapf(err, "could not list %q", resType)
						}
						resources = append(resources, list.GetItems()...)
					}
				}
			}

			switch ctx.args.format {
			case formatUniversal:
				for _, res := range resources {
					if _, err := cmd.OutOrStdout().Write([]byte("---\n")); err != nil {
						return err
					}
					if err := printers.GenericPrint(output.YAMLFormat, res, table.Table{}, cmd.OutOrStdout()); err != nil {
						return err
					}
				}
			case formatKubernetes:
				k8sResources, err := pctx.CurrentKubernetesResourcesClient()
				if err != nil {
					return err
				}
				yamlPrinter := yaml.NewPrinter()
				for _, res := range resources {
					obj, err := k8sResources.Get(cmd.Context(), res.Descriptor(), res.GetMeta().GetName(), res.GetMeta().GetMesh())
					if err != nil {
						return err
					}
					cleanKubeObject(obj)
					if err := yamlPrinter.Print(obj, cmd.OutOrStdout()); err != nil {
						return err
					}
				}
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&ctx.args.profile, "profile", "p", profileFederation, fmt.Sprintf(`Profile. Available values: %q, %q, %q`, profileFederation, profileAll, profileFederationWithPolicies))
	cmd.Flags().StringVarP(&ctx.args.format, "format", "f", formatUniversal, fmt.Sprintf(`Policy format output. Available values: %q, %q`, formatUniversal, formatKubernetes))
	return cmd
}

// cleans kubernetes object, so it can be applied on any other cluster
func cleanKubeObject(obj map[string]interface{}) {
	meta := obj["metadata"].(map[string]interface{})
	delete(meta, "resourceVersion")
	delete(meta, "ownerReferences")
	delete(meta, "uid")
	delete(meta, "generation")
	delete(meta, "managedFields")
}

func resourcesTypesToDump(ctx context.Context, ectx *exportContext) ([]model.ResourceType, error) {
	rt, ok := staticProfiles[ectx.args.profile]
	if ok {
		return rt, nil
	}
	client, err := ectx.CurrentResourcesListClient()
	if err != nil {
		return nil, err
	}
	list, err := client.List(ctx)
	if err != nil {
		return nil, err
	}
	var resTypes []model.ResourceType
	for _, res := range list.Resources {
		switch ectx.args.profile {
		case profileAll:
			resTypes = append(resTypes, model.ResourceType(res.Name))
		case profileFederationWithPolicies:
			if res.IncludeInFederationWithPolicies {
				resTypes = append(resTypes, model.ResourceType(res.Name))
			}
		default:
			return nil, errors.New("invalid profile")
		}
	}
	return resTypes, nil
}
