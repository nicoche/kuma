import{L as D}from"./LoadingBox-d56dce8e.js";import{O as T,a as O,b as B}from"./OnboardingPage-f6a09138.js";import{S as R}from"./StatusBadge-01928c30.js";import{d as S,u as V,m as k,T as F,a as u,o,b as g,w as t,e as s,f as c,c as p,F as x,C as N,t as f,p as h,_ as P}from"./index-203d56a2.js";import{g as C}from"./dataplane-dcd0858b.js";const E={key:0,class:"status-loading-box mb-4"},I={key:1},L={class:"mb-4"},K=S({__name:"DataplanesOverview",setup($){const b=V(),A=[{label:"Mesh",key:"mesh"},{label:"Name",key:"name"},{label:"Status",key:"status"}],a=k({total:0,data:[]}),_=k(null);F(function(){w()}),y();function w(){_.value!==null&&window.clearTimeout(_.value)}async function y(){let i=!1;const m=[];try{const{items:n}=await b.getAllDataplanes({size:10});if(Array.isArray(n)&&n.length>0)for(const v of n){const{name:r,mesh:d}=v,l=await b.getDataplaneOverviewFromMesh({mesh:d,name:r}),e=C(l.dataplaneInsight);e==="offline"&&(i=!0),m.push({status:e,name:r,mesh:d})}else i=!0}catch(n){console.error(n)}a.value.data=m,a.value.total=a.value.data.length,i&&(w(),_.value=window.setTimeout(y,1e3))}return(i,m)=>{const n=u("RouteTitle"),v=u("KTable"),r=u("AppView"),d=u("RouteView");return o(),g(d,{name:"onboarding-dataplanes-overview"},{default:t(({t:l})=>[s(n,{title:l("onboarding.routes.dataplanes-overview.title"),render:!1},null,8,["title"]),c(),s(r,null,{default:t(()=>[s(T,null,{header:t(()=>[(o(!0),p(x,null,N([a.value.data.length>0?"success":"waiting"],e=>(o(),g(O,{key:e,"data-testid":`state-${e}`},{title:t(()=>[c(f(l(`onboarding.routes.dataplanes-overview.header.${e}.title`)),1)]),description:t(()=>[h("p",null,f(l(`onboarding.routes.dataplanes-overview.header.${e}.description`)),1)]),_:2},1032,["data-testid"]))),128))]),content:t(()=>[a.value.data.length===0?(o(),p("div",E,[s(D)])):(o(),p("div",I,[h("p",L,[h("b",null,"Found "+f(a.value.data.length)+" DPPs:",1)]),c(),s(v,{class:"mb-4",fetcher:()=>a.value,headers:A,"disable-pagination":""},{status:t(({rowValue:e})=>[e?(o(),g(R,{key:0,status:e},null,8,["status"])):(o(),p(x,{key:1},[c(`
                  —
                `)],64))]),_:1},8,["fetcher"])]))]),navigation:t(()=>[s(B,{"next-step":"onboarding-completed","previous-step":"onboarding-add-services-code","should-allow-next":a.value.data.length>0},null,8,["should-allow-next"])]),_:2},1024)]),_:2},1024)]),_:1})}}});const q=P(K,[["__scopeId","data-v-9fb2605d"]]);export{q as default};
