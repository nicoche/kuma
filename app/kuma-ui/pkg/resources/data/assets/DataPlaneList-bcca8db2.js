var Ke=Object.defineProperty;var Me=(t,o,a)=>o in t?Ke(t,o,{enumerable:!0,configurable:!0,writable:!0,value:a}):t[o]=a;var X=(t,o,a)=>(Me(t,typeof o!="symbol"?o+"":o,a),a);import{d as de,g as C,cx as Ee,co as De,cy as Oe,cz as Re,cc as qe,o as d,c as y,e as l,b as s,a as E,w as U,bV as T,u as p,j as $,F as N,cd as K,m as ee,bX as ce,bY as pe,k as fe,bP as A,h as Q,c0 as je,cA as xe,cB as Be,ci as Ce,bQ as ze,bR as He,cv as Qe,f as Ge,cm as Pe,cC as Ye,cD as Ze,cr as Je,cs as We,cE as Xe,cq as et,i as H,cF as tt,cG as at,cH as nt,P as Se,bZ as Ae,ce as st,cf as lt}from"./index-08ba2993.js";import{Q as j}from"./QueryParameter-70743f73.js";import{C as ot}from"./ContentWrapper-fe544c43.js";import{D as it}from"./DataOverview-1eb5b106.js";import{T as rt}from"./TagList-e8e9bfa1.js";import{S as ut}from"./StatusBadge-c118c8ba.js";import{_ as dt}from"./YamlView.vue_vue_type_script_setup_true_lang-f673f333.js";import{_ as ct}from"./EmptyBlock.vue_vue_type_script_setup_true_lang-cf69250c.js";const Ie=[{key:"status",label:"Status"},{key:"name",label:"DPP"},{key:"type",label:"Type"},{key:"service",label:"Service"},{key:"protocol",label:"Protocol"},{key:"zone",label:"Zone"},{key:"lastConnected",label:"Last Connected"},{key:"lastUpdated",label:"Last Updated"},{key:"totalUpdates",label:"Total Updates"},{key:"dpVersion",label:"Kuma DP version"},{key:"envoyVersion",label:"Envoy version"},{key:"details",label:"Details",hideLabel:!0}],pt=["name","details"],ft=Ie.filter(t=>!pt.includes(t.key)).map(t=>({tableHeaderKey:t.key,label:t.label,isChecked:!1})),$e=["status","name","type","service","protocol","zone","lastUpdated","dpVersion","details"];function mt(t,o=$e){return Ie.filter(a=>o.includes(a.key)?t?!0:a.key!=="zone":!1)}const vt=["protocol","service","zone"];function gt(t){const o=new Map;for(const[a,w]of t){const i=vt.includes(a),c=i?"tag":a;o.has(c)||o.set(c,[]);const b=o.get(c);let k;c==="tag"?k=(i?`kuma.io/${a}:${w}`:w).replace(/\s+/g,""):k=w,b.push(k.trim())}return o}const L=t=>(ce("data-v-5b05c159"),t=t(),pe(),t),yt={class:"entity-summary entity-section-list"},ht={class:"block-list"},_t={class:"entity-title","data-testid":"data-plane-proxy-title"},bt={class:"definition"},kt=L(()=>l("span",null,"Mesh:",-1)),wt={key:0},Tt=L(()=>l("h4",null,"Tags",-1)),Dt={key:1},Ct=L(()=>l("h4",null,"Dependencies",-1)),Pt={class:"mt-2 heading-with-icon"},St={key:0},At=L(()=>l("h4",null,"Insights",-1)),Et={class:"block-list"},Ot=["data-testid"],xt=L(()=>l("span",null,"Connect time:",-1)),It=["data-testid"],$t=L(()=>l("span",null,"Disconnect time:",-1)),Vt={class:"definition"},Ut=L(()=>l("span",null,"CP instance ID:",-1)),Nt={key:0},Lt=L(()=>l("summary",null,`
              Responses (acknowledged / sent)
            `,-1)),Ft=["data-testid"],Kt={class:"config-section"},Mt=de({__name:"DataPlaneEntitySummary",props:{dataPlaneOverview:{type:Object,required:!0}},setup(t){const o=t,a=C(()=>{const{name:u,mesh:r,dataplane:h}=o.dataPlaneOverview;return{type:"Dataplane",name:u,mesh:r,networking:h.networking}}),w=C(()=>Ee(o.dataPlaneOverview.dataplane)),i=C(()=>{var r;const u=Array.from(((r=o.dataPlaneOverview.dataplaneInsight)==null?void 0:r.subscriptions)??[]);return u.reverse(),u.map(h=>{const f=h.connectTime!==void 0?De(h.connectTime):"—",v=h.disconnectTime!==void 0?De(h.disconnectTime):"—",O=Object.entries(h.status).filter(([D])=>!["total","lastUpdateTime"].includes(D)).map(([D,P])=>{const V=`${P.responsesAcknowledged??0} / ${P.responsesSent??0}`;return{type:D.toUpperCase(),ratio:V,responsesSent:P.responsesSent??0,responsesAcknowledged:P.responsesAcknowledged??0,responsesRejected:P.responsesRejected??0}});return{subscription:h,formattedConnectDate:f,formattedDisconnectDate:v,statuses:O}})}),c=C(()=>{const{status:u}=Oe(o.dataPlaneOverview.dataplane,o.dataPlaneOverview.dataplaneInsight);return u}),b=C(()=>{const u=Re(o.dataPlaneOverview.dataplaneInsight);return u!==null?Object.entries(u).map(([r,h])=>({name:r,version:h})):[]}),k=C(()=>{var P;const u=((P=o.dataPlaneOverview.dataplaneInsight)==null?void 0:P.subscriptions)??[];if(u.length===0)return[];const r=u[u.length-1];if(!r.version)return[];const h=[],f=r.version.envoy,v=r.version.kumaDp;if(!(f.kumaDpCompatible!==void 0?f.kumaDpCompatible:!0)){const V=`Envoy ${f.version} is not supported by Kuma DP ${v.version}.`;h.push(V)}if(!(v.kumaCpCompatible!==void 0?v.kumaCpCompatible:!0)){const V=`Kuma DP ${v.version} is not supported by this Kuma control plane.`;h.push(V)}return h});return(u,r)=>{const h=qe("router-link");return d(),y("div",yt,[l("section",null,[l("div",ht,[l("div",null,[l("h3",_t,[l("span",null,[s(`
              DPP:

              `),E(h,{to:{name:"data-plane-detail-view",params:{mesh:t.dataPlaneOverview.mesh,dataPlane:t.dataPlaneOverview.name}}},{default:U(()=>[s(T(t.dataPlaneOverview.name),1)]),_:1},8,["to"])]),s(),E(ut,{status:p(c)},null,8,["status"])]),s(),l("div",bt,[kt,s(),l("span",null,T(t.dataPlaneOverview.mesh),1)])]),s(),p(w).length>0?(d(),y("div",wt,[Tt,s(),E(rt,{tags:p(w)},null,8,["tags"])])):$("",!0),s(),p(b).length>0?(d(),y("div",Dt,[Ct,s(),(d(!0),y(N,null,K(p(b),(f,v)=>(d(),y("div",{key:v,class:"definition"},[l("span",null,T(f.name)+":",1),s(),l("span",null,T(f.version),1)]))),128)),s(),p(k).length>0?(d(),y(N,{key:0},[l("h5",Pt,[s(`
              Warnings

              `),E(p(ee),{class:"ml-1",icon:"warning",color:"var(--black-75)","secondary-color":"var(--yellow-300)",size:"20"})]),s(),(d(!0),y(N,null,K(p(k),(f,v)=>(d(),y("p",{key:v},T(f),1))),128))],64)):$("",!0)])):$("",!0)])]),s(),p(i).length>0?(d(),y("section",St,[At,s(),l("div",Et,[(d(!0),y(N,null,K(p(i),(f,v)=>(d(),y("div",{key:v},[l("div",{class:"definition","data-testid":`data-plane-connect-time-${v}`},[xt,s(),l("span",null,T(f.formattedConnectDate),1)],8,Ot),s(),l("div",{class:"definition","data-testid":`data-plane-disconnect-time-${v}`},[$t,s(),l("span",null,T(f.formattedDisconnectDate),1)],8,It),s(),l("div",Vt,[Ut,s(),l("span",null,T(f.subscription.controlPlaneInstanceId),1)]),s(),f.statuses.length>0?(d(),y("details",Nt,[Lt,s(),(d(!0),y(N,null,K(f.statuses,(O,D)=>(d(),y("div",{key:`${v}-${D}`,class:"definition","data-testid":`data-plane-subscription-status-${v}-${D}`},[l("span",null,T(O.type)+":",1),s(),l("span",null,T(O.ratio),1)],8,Ft))),128))])):$("",!0)]))),128))])])):$("",!0),s(),l("section",Kt,[E(dt,{id:"code-block-data-plane-summary",content:p(a),"code-max-height":"250px"},null,8,["content"])])])}}});const Rt=fe(Mt,[["__scopeId","data-v-5b05c159"]]);function qt(t,o,a){return Math.max(o,Math.min(t,a))}const jt=["ControlLeft","ControlRight","ShiftLeft","ShiftRight","AltLeft"];class Bt{constructor(o,a){X(this,"commands");X(this,"keyMap");X(this,"boundTriggerShortcuts");this.commands=a,this.keyMap=Object.fromEntries(Object.entries(o).map(([w,i])=>[w.toLowerCase(),i])),this.boundTriggerShortcuts=this.triggerShortcuts.bind(this)}registerListener(){document.addEventListener("keydown",this.boundTriggerShortcuts)}unRegisterListener(){document.removeEventListener("keydown",this.boundTriggerShortcuts)}triggerShortcuts(o){zt(o,this.keyMap,this.commands)}}function zt(t,o,a){const w=Ht(t.code),i=[t.ctrlKey?"ctrl":"",t.shiftKey?"shift":"",t.altKey?"alt":"",w].filter(k=>k!=="").join("+"),c=o[i];if(!c)return;const b=a[c];b.isAllowedContext&&!b.isAllowedContext(t)||(b.shouldPreventDefaultAction&&t.preventDefault(),!(b.isDisabled&&b.isDisabled())&&b.trigger(t))}function Ht(t){return jt.includes(t)?"":t.replace(/^Key/,"").toLowerCase()}function Qt(t,o){const a=" "+t,w=a.matchAll(/ ([-\s\w]+):\s*/g),i=[];for(const c of Array.from(w)){if(c.index===void 0)continue;const b=Gt(c[1]);if(o.length>0&&!o.includes(b))throw new Error(`Unknown field “${b}”. Known fields: ${o.join(", ")}`);const k=c.index+c[0].length,u=a.substring(k);let r;if(/^\s*["']/.test(u)){const f=u.match(/['"](.*?)['"]/);if(f!==null)r=f[1];else throw new Error(`Quote mismatch for field “${b}”.`)}else{const f=u.indexOf(" "),v=f===-1?u.length:f;r=u.substring(0,v)}r!==""&&i.push([b,r])}return i}function Gt(t){return t.trim().replace(/\s+/g,"-").replace(/-[a-z]/g,(o,a)=>a===0?o:o.substring(1).toUpperCase())}const Ve=t=>(ce("data-v-e4356188"),t=t(),pe(),t),Yt=Ve(()=>l("span",{class:"visually-hidden"},"Focus filter",-1)),Zt=["for"],Jt=["id","placeholder"],Wt={key:0,class:"k-suggestion-box","data-testid":"k-filter-bar-suggestion-box"},Xt={class:"k-suggestion-list"},ea={key:0,class:"k-filter-bar-error"},ta={key:0},aa=["title","data-filter-field"],na={class:"visually-hidden"},sa=Ve(()=>l("span",{class:"visually-hidden"},"Clear query",-1)),la=de({__name:"KFilterBar",props:{id:{type:String,required:!0},fields:{type:Object,required:!0},placeholder:{type:String,required:!1,default:null},query:{type:String,required:!1,default:""}},emits:["fields-change"],setup(t,{emit:o}){const a=t,w=A(null),i=A(null),c=A(a.query),b=A([]),k=A(null),u=A(!1),r=A(-1),h=C(()=>Object.keys(a.fields)),f=C(()=>Object.entries(a.fields).slice(0,5).map(([e,n])=>({fieldName:e,...n}))),v=C(()=>h.value.length>0?`Filter by ${h.value.join(", ")}`:"Filter"),O=C(()=>a.placeholder??v.value);Q(()=>b.value,function(e,n){oe(e,n)||(k.value=null,o("fields-change",{fields:e,query:c.value}))}),Q(()=>c.value,function(){c.value===""&&(k.value=null),u.value=!0});const D={Enter:"submitQuery",Escape:"closeSuggestionBox",ArrowDown:"jumpToNextSuggestion",ArrowUp:"jumpToPreviousSuggestion"},P={submitQuery:{trigger:B,isAllowedContext(e){return i.value!==null&&e.composedPath().includes(i.value)},shouldPreventDefaultAction:!0},jumpToNextSuggestion:{trigger:ae,isAllowedContext(e){return i.value!==null&&e.composedPath().includes(i.value)},shouldPreventDefaultAction:!0},jumpToPreviousSuggestion:{trigger:ne,isAllowedContext(e){return i.value!==null&&e.composedPath().includes(i.value)},shouldPreventDefaultAction:!0},closeSuggestionBox:{trigger:F,isAllowedContext(e){return w.value!==null&&e.composedPath().includes(w.value)}}};function V(){const e=new Bt(D,P);ze(function(){e.registerListener()}),He(function(){e.unRegisterListener()}),M(c.value)}V();function te(e){const n=e.target;M(n.value)}function B(){if(i.value instanceof HTMLInputElement)if(r.value===-1)M(i.value.value),u.value=!1;else{const e=f.value[r.value].fieldName;e&&Y(i.value,e)}}function ae(){G(1)}function ne(){G(-1)}function G(e){r.value=qt(r.value+e,-1,f.value.length-1)}function z(){i.value instanceof HTMLInputElement&&i.value.focus()}function se(e){const m=e.currentTarget.getAttribute("data-filter-field");m&&i.value instanceof HTMLInputElement&&Y(i.value,m)}function Y(e,n){const m=c.value===""||c.value.endsWith(" ")?"":" ";c.value+=m+n+":",e.focus(),r.value=-1}function le(){c.value="",i.value instanceof HTMLInputElement&&(i.value.value="",i.value.focus(),M(""))}function Z(e){e.relatedTarget===null&&F(),w.value instanceof HTMLElement&&e.relatedTarget instanceof Node&&!w.value.contains(e.relatedTarget)&&F()}function F(){u.value=!1}function M(e){k.value=null;try{const n=Qt(e,h.value);n.sort((m,_)=>m[0].localeCompare(_[0])),b.value=n}catch(n){if(n instanceof Error)k.value=n,u.value=!0;else throw n}}function oe(e,n){return JSON.stringify(e)===JSON.stringify(n)}return(e,n)=>(d(),y("div",{ref_key:"filterBar",ref:w,class:"k-filter-bar","data-testid":"k-filter-bar"},[l("button",{class:"k-focus-filter-input-button",title:"Focus filter",type:"button","data-testid":"k-filter-bar-focus-filter-input-button",onClick:z},[Yt,s(),E(p(ee),{"aria-hidden":"true",class:"k-filter-icon",color:"var(--grey-400)","data-testid":"k-filter-bar-filter-icon","hide-title":"",icon:"filter",size:"20"})]),s(),l("label",{for:`${a.id}-filter-bar-input`,class:"visually-hidden"},[je(e.$slots,"default",{},()=>[s(T(p(v)),1)],!0)],8,Zt),s(),xe(l("input",{id:`${a.id}-filter-bar-input`,ref_key:"filterInput",ref:i,"onUpdate:modelValue":n[0]||(n[0]=m=>c.value=m),class:"k-filter-bar-input",type:"text",placeholder:p(O),"data-testid":"k-filter-bar-filter-input",onFocus:n[1]||(n[1]=m=>u.value=!0),onBlur:Z,onChange:te},null,40,Jt),[[Be,c.value]]),s(),u.value?(d(),y("div",Wt,[l("div",Xt,[k.value!==null?(d(),y("p",ea,T(k.value.message),1)):(d(),y("button",{key:1,class:Ce(["k-submit-query-button",{"k-submit-query-button-is-selected":r.value===-1}]),title:"Submit query",type:"button","data-testid":"k-filter-bar-submit-query-button",onClick:B},`
          Submit `+T(c.value),3)),s(),(d(!0),y(N,null,K(p(f),(m,_)=>(d(),y("div",{key:`${a.id}-${_}`,class:Ce(["k-suggestion-list-item",{"k-suggestion-list-item-is-selected":r.value===_}])},[l("b",null,T(m.fieldName),1),m.description!==""?(d(),y("span",ta,": "+T(m.description),1)):$("",!0),s(),l("button",{class:"k-apply-suggestion-button",title:`Add ${m.fieldName}:`,type:"button","data-filter-field":m.fieldName,"data-testid":"k-filter-bar-apply-suggestion-button",onClick:se},[l("span",na,"Add "+T(m.fieldName)+":",1),s(),E(p(ee),{"aria-hidden":"true",color:"currentColor","hide-title":"",icon:"chevronRight",size:"16"})],8,aa)],2))),128))])])):$("",!0),s(),c.value!==""?(d(),y("button",{key:1,class:"k-clear-query-button",title:"Clear query",type:"button","data-testid":"k-filter-bar-clear-query-button",onClick:le},[sa,s(),E(p(ee),{"aria-hidden":"true",color:"currentColor",icon:"clear","hide-title":"",size:"20"})])):$("",!0)],512))}});const oa=fe(la,[["__scopeId","data-v-e4356188"]]),ia=t=>(ce("data-v-3f7c3086"),t=t(),pe(),t),ra={key:0},ua=ia(()=>l("label",{for:"data-planes-type-filter",class:"mr-2"},`
              Type:
            `,-1)),da=["value"],ca=["for"],pa=["id","checked","onChange"],fa=de({__name:"DataPlaneList",props:{dataPlaneOverviews:{type:Array,required:!0},isLoading:{type:Boolean,required:!1,default:!1},error:{type:Error,required:!1,default:null},nextUrl:{type:String,required:!1,default:null},pageOffset:{type:Number,required:!1,default:0},selectedDppName:{type:String,required:!1,default:null},isGatewayView:{type:Boolean,required:!1,default:!1},dppFilterFields:{type:Object,required:!0}},emits:["load-data"],setup(t,{emit:o}){const a=t,w=50,i={All:"true",Builtin:"builtin",Delegated:"delegated"},c=["All","Builtin","Delegated"],b={title:"No Data",message:"There are no data plane proxies present."},k=Qe(),u=Ge(),r=A($e),h=A({headers:[],data:[]}),f=A(j.get("filterQuery")??""),v=A("All"),O=A({}),D=A(null),P=C(()=>u.getters["config/getMulticlusterStatus"]),V=C(()=>({name:u.getters["config/getEnvironment"]==="universal"?"universal-dataplane":"kubernetes-dataplane"})),te=C(()=>"tag"in a.dppFilterFields?'tag: "kuma.io/protocol: http"':"name"in a.dppFilterFields?"name: cluster":"field: value"),B=C(()=>{let e=mt(P.value,r.value);return a.isGatewayView?e=e.filter(n=>n.key!=="protocol"):e=e.filter(n=>n.key!=="type"),{data:h.value.data,headers:e}}),ae=C(()=>ft.filter(e=>a.isGatewayView?e.tableHeaderKey!=="protocol":e.tableHeaderKey!=="type").filter(e=>P.value?!0:e.tableHeaderKey!=="zone").map(e=>{const n=r.value.includes(e.tableHeaderKey);return{...e,isChecked:n}}));Q(v,function(){z(0)}),Q(O,function(){z(0)}),Q(()=>a.dataPlaneOverviews,function(){Z()});function ne(){const e=Ae.get("dpVisibleTableHeaderKeys");Array.isArray(e)&&(r.value=e),Z()}ne();function G(e){z(e)}function z(e){const n={...O.value};"gateway"in n||(n.gateway=i[v.value]),o("load-data",e,n)}function se(e){e.stopPropagation()}function Y(e,n){const m=e.target,_=r.value.findIndex(x=>x===n);m.checked&&_===-1?r.value.push(n):!m.checked&&_>-1&&r.value.splice(_,1),Ae.set("dpVisibleTableHeaderKeys",Array.from(new Set(r.value)))}function le(){st.logger.info(lt.CREATE_DATA_PLANE_PROXY_CLICKED)}async function Z(){try{Array.isArray(a.dataPlaneOverviews)&&a.dataPlaneOverviews.length>0?(F(a.selectedDppName??a.dataPlaneOverviews[0].name),h.value.data=await Promise.all(a.dataPlaneOverviews.map(e=>M(e)))):(F(null),h.value.data=[])}catch(e){console.error(e)}}function F(e){e&&a.dataPlaneOverviews.length>0?(D.value=a.dataPlaneOverviews.find(n=>n.name===e)??a.dataPlaneOverviews[0],j.set(a.isGatewayView?"gateway":"dpp",D.value.name)):(D.value=null,j.set(a.isGatewayView?"gateway":"dpp",null))}async function M(e){var ge,ye,he,_e,be;const n=e.mesh,m=e.name,_=((ge=e.dataplane.networking.gateway)==null?void 0:ge.type)||"STANDARD",x={name:_==="STANDARD"?"data-plane-detail-view":"gateway-detail-view",params:{mesh:n,dataPlane:m}},R={name:"mesh-detail-view",params:{mesh:n}},J=["kuma.io/protocol","kuma.io/service","kuma.io/zone"],W=Ee(e.dataplane).filter(g=>J.includes(g.label)),ie=(ye=W.find(g=>g.label==="kuma.io/service"))==null?void 0:ye.value,Ue=(he=W.find(g=>g.label==="kuma.io/protocol"))==null?void 0:he.value,re=(_e=W.find(g=>g.label==="kuma.io/zone"))==null?void 0:_e.value;let me;ie!==void 0&&(me={name:"service-detail-view",params:{mesh:n,service:ie}});let ve;re!==void 0&&(ve={name:"zones",query:{ns:re}});const{status:Ne}=Oe(e.dataplane,e.dataplaneInsight),Le=((be=e.dataplaneInsight)==null?void 0:be.subscriptions)??[],Fe={totalUpdates:0,totalRejectedUpdates:0,dpVersion:null,envoyVersion:null,selectedTime:NaN,selectedUpdateTime:NaN,version:null},S=Le.reduce((g,I)=>{var ke,we;if(I.connectTime){const Te=Date.parse(I.connectTime);(!g.selectedTime||Te>g.selectedTime)&&(g.selectedTime=Te)}const ue=Date.parse(I.status.lastUpdateTime);return ue&&(!g.selectedUpdateTime||ue>g.selectedUpdateTime)&&(g.selectedUpdateTime=ue),{totalUpdates:g.totalUpdates+parseInt(I.status.total.responsesSent??"0",10),totalRejectedUpdates:g.totalRejectedUpdates+parseInt(I.status.total.responsesRejected??"0",10),dpVersion:((ke=I.version)==null?void 0:ke.kumaDp.version)||g.dpVersion,envoyVersion:((we=I.version)==null?void 0:we.envoy.version)||g.envoyVersion,selectedTime:g.selectedTime,selectedUpdateTime:g.selectedUpdateTime,version:I.version||g.version}},Fe),q={name:m,nameRoute:x,mesh:n,meshRoute:R,type:_,zone:re??"—",zoneRoute:ve,service:ie??"—",serviceInsightRoute:me,protocol:Ue??"—",status:Ne,totalUpdates:S.totalUpdates,totalRejectedUpdates:S.totalRejectedUpdates,dpVersion:S.dpVersion??"—",envoyVersion:S.envoyVersion??"—",warnings:[],unsupportedEnvoyVersion:!1,unsupportedKumaDPVersion:!1,kumaDpAndKumaCpMismatch:!1,lastUpdated:S.selectedUpdateTime?Pe(new Date(S.selectedUpdateTime).toUTCString()):"—",lastConnected:S.selectedTime?Pe(new Date(S.selectedTime).toUTCString()):"—",overview:e};if(S.version){const{kind:g}=Ye(S.version);switch(g!==Ze&&q.warnings.push(g),g){case We:q.unsupportedEnvoyVersion=!0;break;case Je:q.unsupportedKumaDPVersion=!0;break}}return P.value&&S.dpVersion&&W.find(I=>I.label===Xe)&&typeof S.version.kumaDp.kumaCpCompatible=="boolean"&&!S.version.kumaDp.kumaCpCompatible&&(q.warnings.push(et),q.kumaDpAndKumaCpMismatch=!0),q}function oe({fields:e,query:n}){const m=j.get("filterFields"),_=m!==null?JSON.parse(m):{},x=JSON.stringify(_),R=Object.fromEntries(gt(e)),J=JSON.stringify(R);j.set("filterQuery",n||null),j.set("filterFields",J),x!==J&&(O.value=R)}return(e,n)=>(d(),H(ot,null,{content:U(()=>{var m;return[E(it,{"selected-entity-name":(m=D.value)==null?void 0:m.name,"page-size":w,"is-loading":a.isLoading,error:t.error,"empty-state":b,"table-data":p(B),"table-data-is-empty":p(B).data.length===0,"show-details":"",next:a.nextUrl!==null,"page-offset":a.pageOffset,onTableAction:n[1]||(n[1]=_=>F(_.name)),onLoadData:G},{additionalControls:U(()=>[E(oa,{id:"data-plane-proxy-filter",class:"data-plane-proxy-filter",placeholder:p(te),query:f.value,fields:a.dppFilterFields,onFieldsChange:oe},null,8,["placeholder","query","fields"]),s(),a.isGatewayView?(d(),y("div",ra,[ua,s(),xe(l("select",{id:"data-planes-type-filter","onUpdate:modelValue":n[0]||(n[0]=_=>v.value=_),"data-testid":"data-planes-type-filter"},[(d(!0),y(N,null,K(p(c),(_,x)=>(d(),y("option",{key:x,value:_},T(_),9,da))),128))],512),[[tt,v.value]])])):$("",!0),s(),E(p(at),{label:"Columns",icon:"cogwheel","button-appearance":"outline"},{items:U(()=>[l("div",{onClick:se},[(d(!0),y(N,null,K(p(ae),(_,x)=>(d(),H(p(nt),{key:x,class:"table-header-selector-item",item:_},{default:U(()=>[l("label",{for:`data-plane-table-header-checkbox-${x}`,class:"k-checkbox table-header-selector-item-checkbox"},[l("input",{id:`data-plane-table-header-checkbox-${x}`,checked:_.isChecked,type:"checkbox",class:"k-input",onChange:R=>Y(R,_.tableHeaderKey)},null,40,pa),s(" "+T(_.label),1)],8,ca)]),_:2},1032,["item"]))),128))])]),_:1}),s(),E(p(Se),{appearance:"creation",to:p(V),icon:"plus","data-testid":"data-plane-create-data-plane-button",onClick:le},{default:U(()=>[s(`
            Create data plane proxy
          `)]),_:1},8,["to"]),s(),p(k).query.ns?(d(),H(p(Se),{key:1,appearance:"primary",icon:"arrowLeft",to:{name:p(k).name},"data-testid":"data-plane-ns-back-button"},{default:U(()=>[s(`
            View All
          `)]),_:1},8,["to"])):$("",!0)]),_:1},8,["selected-entity-name","is-loading","error","table-data","table-data-is-empty","next","page-offset"])]}),sidebar:U(()=>[D.value!==null?(d(),H(Rt,{key:0,"data-plane-overview":D.value},null,8,["data-plane-overview"])):(d(),H(ct,{key:1}))]),_:1}))}});const Ta=fe(fa,[["__scopeId","data-v-3f7c3086"]]);export{Ta as D};
