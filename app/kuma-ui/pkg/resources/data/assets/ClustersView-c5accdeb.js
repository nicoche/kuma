import{E as l}from"./EnvoyData-4c52d667.js";import{g as d}from"./dataplane-dcd0858b.js";import{d as _,a as e,o as g,b as h,w as t,e as o,p as f,f as w,q as C}from"./index-203d56a2.js";import"./index-9dd3e7d3.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-f7a0d7a8.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-01997eab.js";import"./ErrorBlock-085322b0.js";import"./TextWithCopyButton-45b0690a.js";import"./CopyButton-4a565fd0.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-297a38e2.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-096f0e9b.js";const b=_({__name:"ClustersView",props:{data:{}},setup(r){const a=r;return(V,z)=>{const n=e("RouteTitle"),p=e("KCard"),c=e("AppView"),i=e("RouteView");return g(),h(i,{name:"zone-egress-clusters-view",params:{zoneEgress:"",codeSearch:""}},{default:t(({route:s,t:m})=>[o(c,null,{title:t(()=>[f("h2",null,[o(n,{title:m("zone-egresses.routes.item.navigation.zone-egress-clusters-view")},null,8,["title"])])]),default:t(()=>[w(),o(p,null,{body:t(()=>[o(l,{status:C(d)(a.data.zoneEgressInsight),resource:"Zone",src:`/zone-egresses/${s.params.zoneEgress}/data-path/clusters`,query:s.params.codeSearch,onQueryChange:u=>s.update({codeSearch:u})},null,8,["status","src","query","onQueryChange"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{b as default};
