import{K as b}from"./index-9dd3e7d3.js";import{d as f,a as s,o as k,b as g,w as e,e as t,p as y,f as a,t as i,q as r,aA as K}from"./index-203d56a2.js";import{A as C}from"./AppCollection-0aeb195e.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang-01997eab.js";const S=f({__name:"ServicesView",props:{data:{}},setup(l){const m=l;return(V,z)=>{const p=s("RouteTitle"),c=s("RouterLink"),u=s("KButton"),d=s("KDropdownItem"),_=s("KDropdownMenu"),v=s("KCard"),w=s("AppView"),h=s("RouteView");return k(),g(h,{name:"zone-ingress-services-view"},{default:e(({t:n})=>[t(w,null,{title:e(()=>[y("h2",null,[t(p,{title:n("zone-ingresses.routes.item.navigation.zone-ingress-services-view")},null,8,["title"])])]),default:e(()=>[a(),t(v,null,{body:e(()=>[t(C,{"data-testid":"available-services-collection","empty-state-message":n("common.emptyState.message",{type:"Services"}),headers:[{label:"Name",key:"name"},{label:"Mesh",key:"mesh"},{label:"Protocol",key:"protocol"},{label:"No. Instances",key:"instances"},{label:"Actions",key:"actions",hideLabel:!0}],items:m.data.zoneIngress.availableServices},{name:e(({row:o})=>[t(c,{to:{name:"service-detail-view",params:{mesh:o.mesh,service:o.tags["kuma.io/service"]}}},{default:e(()=>[a(i(o.tags["kuma.io/service"]),1)]),_:2},1032,["to"])]),mesh:e(({row:o})=>[t(c,{to:{name:"mesh-detail-view",params:{mesh:o.mesh}}},{default:e(()=>[a(i(o.mesh),1)]),_:2},1032,["to"])]),protocol:e(({row:o})=>[a(i(o.tags["kuma.io/protocol"]??n("common.collection.none")),1)]),instances:e(({row:o})=>[a(i(o.instances),1)]),actions:e(({row:o})=>[t(_,{class:"actions-dropdown","kpop-attributes":{placement:"bottomEnd",popoverClasses:"mt-5 more-actions-popover"},width:"150"},{default:e(()=>[t(u,{class:"non-visual-button",appearance:"secondary",size:"small"},{default:e(()=>[t(r(K),{size:r(b)},null,8,["size"])]),_:1})]),items:e(()=>[t(d,{item:{to:{name:"service-detail-view",params:{mesh:o.mesh,service:o.tags["kuma.io/service"]}},label:n("common.collection.actions.view")}},null,8,["item"])]),_:2},1024)]),_:2},1032,["empty-state-message","headers","items"])]),_:2},1024)]),_:2},1024)]),_:1})}}});export{S as default};
