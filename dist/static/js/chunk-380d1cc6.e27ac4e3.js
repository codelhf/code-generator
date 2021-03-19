(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-380d1cc6"],{1307:function(e,t,a){"use strict";var r=a("15f5"),i=a.n(r);i.a},"15f5":function(e,t,a){},"24d2":function(e,t,a){"use strict";a.d(t,"d",(function(){return i})),a.d(t,"e",(function(){return n})),a.d(t,"b",(function(){return l})),a.d(t,"f",(function(){return o})),a.d(t,"a",(function(){return c})),a.d(t,"c",(function(){return s}));var r=a("b775");function i(e){return Object(r["a"])({url:"/project/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/project/select",method:"get",params:{id:e}})}function l(e){return Object(r["a"])({url:"/project/insert",method:"post",data:e})}function o(e){return Object(r["a"])({url:"/project/update",method:"put",data:e})}function c(e){return Object(r["a"])({url:"/project/delete",method:"delete",params:{ids:e}})}function s(){return Object(r["a"])({url:"/project/lastId",method:"get"})}},8128:function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"page-container"},[a("table-list",{attrs:{data:e.tableList,"show-icon":!0},on:{itemClick:e.handleTableClick}},[a("template-list",{key:e.projectId,ref:"templateList",attrs:{"project-id":e.projectId}}),a("el-form",{ref:"currentTable",staticStyle:{"margin-top":"20px"},attrs:{inline:!0,"label-width":"80px"},model:{value:e.currentTable,callback:function(t){e.currentTable=t},expression:"currentTable"}},[a("el-form-item",{attrs:{label:e.$t("projectGenerate.currentTable.domainName")}},[a("el-input",{attrs:{placeholder:e.$t("projectGenerate.currentTable.placeholderDomainName")},model:{value:e.currentTable.domainName,callback:function(t){e.$set(e.currentTable,"domainName",t)},expression:"currentTable.domainName"}})],1),a("el-form-item",{attrs:{label:e.$t("projectGenerate.currentTable.domainDesc")}},[a("el-input",{attrs:{placeholder:e.$t("projectGenerate.currentTable.placeholderDomainDesc")},model:{value:e.currentTable.domainDesc,callback:function(t){e.$set(e.currentTable,"domainDesc",t)},expression:"currentTable.domainDesc"}})],1),a("el-form-item",{attrs:{label:" "}},[a("el-button",{attrs:{type:"primary",icon:"el-icon-video-play",loading:e.loading},on:{click:e.handleGenerate}},[e._v(e._s(e.$t("projectGenerate.button.generateCode")))]),a("el-button",{attrs:{type:"primary",icon:"el-icon-plus"},on:{click:e.handleDetail}},[e._v(e._s(e.$t("projectGenerate.button.addTemplate")))])],1)],1),a("div",{staticStyle:{width:"100%",height:"calc(((100vh - 124px) / 2) - 82px)"}},[a("el-input",{attrs:{type:"textarea",rows:10},model:{value:e.generateResult,callback:function(t){e.generateResult=t},expression:"generateResult"}})],1)],1)],1)},i=[],n=(a("d81d"),a("b0c0"),a("24d2")),l=a("b775");function o(e){return Object(l["a"])({url:"/projectTable/all",method:"get",params:e})}function c(e,t){return Object(l["a"])({url:"/projectTable/select",method:"get",params:{projectId:e,tableName:t}})}function s(e,t){return Object(l["a"])({url:"/projectTable/generate?projectId="+e,method:"post",data:t})}var u=a("a6fb"),p=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticStyle:{width:"100%",height:"calc((100vh - 124px) / 2)"}},[a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],key:e.tableKey,attrs:{data:e.list.filter((function(t){return!e.search||t.name.toLowerCase().includes(e.search.toLowerCase())})),height:"100%","highlight-current-row":"",border:"border"}},[a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.name"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.name))])]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.directory"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.directory))])]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.packageName"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.packageName))])]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.fileSuffix"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.fileSuffix))])]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.isGenerate"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-switch",{attrs:{"active-value":1,"inactive-value":2,"active-text":e.$t("projectTemplate.switch.activeText"),"inactive-text":e.$t("projectTemplate.switch.inactiveText"),"active-color":"#13ce66","inactive-color":"#ff4949"},on:{change:function(a){return e.handleTableSwitch(t.row)}},model:{value:t.row.isGenerate,callback:function(a){e.$set(t.row,"isGenerate",a)},expression:"scope.row.isGenerate"}})]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.isOverride"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-switch",{attrs:{"active-value":1,"inactive-value":2,"active-text":e.$t("projectTemplate.switch.activeText"),"inactive-text":e.$t("projectTemplate.switch.inactiveText"),"active-color":"#13ce66","inactive-color":"#ff4949"},on:{change:function(a){return e.handleTableSwitch(t.row)}},model:{value:t.row.isOverride,callback:function(a){e.$set(t.row,"isOverride",a)},expression:"scope.row.isOverride"}})]}}])}),a("el-table-column",{attrs:{label:e.$t("projectTemplate.table.operation"),align:"center"},scopedSlots:e._u([{key:"header",fn:function(t){return[a("el-input",{attrs:{size:"mini",placeholder:e.$t("projectTemplate.table.operation")},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})]}},{key:"default",fn:function(t){return[a("el-button",{attrs:{circle:"",size:"mini",icon:"el-icon-edit"},on:{click:function(a){return e.handleDetail(t.row.id)}}}),a("el-button",{attrs:{circle:"",size:"mini",icon:"el-icon-delete"},on:{click:function(a){return e.handleDelete(t.row.id)}}})]}}])})],1),a("el-dialog",{attrs:{title:e.projectTemplate.id?e.$t("projectTemplate.item.editTitle"):e.$t("projectTemplate.item.addTitle"),visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t},close:function(t){return e.handleFormClose("projectTemplateForm")}}},[a("el-form",{ref:"projectTemplateForm",attrs:{model:e.projectTemplate,rules:e.projectTemplateRules(),"label-width":"130px","label-suffix":":"}},[a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.groupId"),prop:"groupId"}},[a("el-select",{attrs:{filterable:"",placeholder:e.$t("projectTemplate.item.placeholderGroup")},on:{change:e.handleGroup},model:{value:e.projectTemplate.groupId,callback:function(t){e.$set(e.projectTemplate,"groupId",t)},expression:"projectTemplate.groupId"}},e._l(e.groupList,(function(e){return a("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1)],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.templateId"),prop:"templateId"}},[a("el-select",{attrs:{filterable:"",placeholder:e.$t("projectTemplate.item.placeholderName")},model:{value:e.projectTemplate.templateId,callback:function(t){e.$set(e.projectTemplate,"templateId",t)},expression:"projectTemplate.templateId"}},e._l(e.templateList,(function(e){return a("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1)],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.directory"),prop:"directory"}},[a("el-input",{attrs:{placeholder:e.$t("projectTemplate.item.placeholderDirectory")},model:{value:e.projectTemplate.directory,callback:function(t){e.$set(e.projectTemplate,"directory",t)},expression:"projectTemplate.directory"}},[a("el-button",{attrs:{slot:"append",icon:"el-icon-folder",type:"success"},on:{click:function(e){}},slot:"append"})],1)],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.packageName"),prop:"packageName"}},[a("el-input",{attrs:{placeholder:e.$t("projectTemplate.item.placeholderPackageName")},model:{value:e.projectTemplate.packageName,callback:function(t){e.$set(e.projectTemplate,"packageName",t)},expression:"projectTemplate.packageName"}})],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.fileSuffix"),prop:"fileSuffix"}},[a("el-input",{attrs:{placeholder:e.$t("projectTemplate.item.placeholderFileSuffix")},model:{value:e.projectTemplate.fileSuffix,callback:function(t){e.$set(e.projectTemplate,"fileSuffix",t)},expression:"projectTemplate.fileSuffix"}})],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.isGenerate"),prop:"isGenerate"}},[a("el-switch",{attrs:{"active-value":1,"inactive-value":2,"active-text":e.$t("projectTemplate.switch.activeText"),"inactive-text":e.$t("projectTemplate.switch.inactiveText"),"active-color":"#13ce66","inactive-color":"#ff4949"},model:{value:e.projectTemplate.isGenerate,callback:function(t){e.$set(e.projectTemplate,"isGenerate",t)},expression:"projectTemplate.isGenerate"}})],1),a("el-form-item",{attrs:{label:e.$t("projectTemplate.item.isOverride"),prop:"isOverride"}},[a("el-switch",{attrs:{"active-value":1,"inactive-value":2,"active-text":e.$t("projectTemplate.switch.activeText"),"inactive-text":e.$t("projectTemplate.switch.inactiveText"),"active-color":"#13ce66","inactive-color":"#ff4949"},model:{value:e.projectTemplate.isOverride,callback:function(t){e.$set(e.projectTemplate,"isOverride",t)},expression:"projectTemplate.isOverride"}})],1)],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){return e.handleFormClose("projectTemplateForm")}}},[e._v(e._s(e.$t("common.form.cancel")))]),a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleFormSubmit("projectTemplateForm")}}},[e._v(e._s(e.$t("common.form.confirm")))])],1)],1)],1)},m=[];a("4de4");function d(e){return Object(l["a"])({url:"/projectTemplate/list",method:"get",params:e})}function f(e){return Object(l["a"])({url:"/projectTemplate/select",method:"get",params:{id:e}})}function h(e){return Object(l["a"])({url:"/projectTemplate/insert",method:"post",data:e})}function b(e){return Object(l["a"])({url:"/projectTemplate/update",method:"put",data:e})}function j(e){return Object(l["a"])({url:"/projectTemplate/delete",method:"delete",params:{ids:e}})}var g=a("d9f8"),T=a("c621"),v={name:"ProjectTemplate",props:{projectId:{type:String,required:!0}},data:function(){return{tableKey:0,list:[],listLoading:!0,search:"",groupList:[],templateAll:[],templateList:[],dialogFormVisible:!1,projectTemplate:{id:"",projectId:"",groupId:"",templateId:"",directory:"",packageName:"",fileSuffix:"",isGenerate:1,isOverride:1,name:"",type:null,fileType:""}}},created:function(){this.getList(),this.getGroupList(),this.getTemplateList()},methods:{projectTemplateRules:function(){return{projectId:[{required:!0,message:this.$t("projectTemplate.itemRules.projectId"),trigger:"blur"}],groupId:[{required:!0,message:this.$t("projectTemplate.itemRules.groupId"),trigger:"blur"}],templateId:[{required:!0,message:this.$t("projectTemplate.itemRules.templateId"),trigger:"blur"}],directory:[{required:!0,message:this.$t("projectTemplate.itemRules.directory"),trigger:"blur"}],packageName:[{required:!0,message:this.$t("projectTemplate.itemRules.packageName"),trigger:"blur"}],fileSuffix:[{required:!0,message:this.$t("projectTemplate.itemRules.fileSuffix"),trigger:"blur"}],isGenerate:[{required:!0,message:this.$t("projectTemplate.itemRules.isGenerate"),trigger:"blur"}],isOverride:[{required:!0,message:this.$t("projectTemplate.itemRules.isOverride"),trigger:"blur"}]}},getList:function(){var e=this,t={projectId:this.projectId};d(t).then((function(t){e.list=t.data.list,e.total=t.data.total,setTimeout((function(){e.listLoading=!1}),1500)}))},getGroupList:function(){var e=this;Object(g["d"])({}).then((function(t){e.groupList=t.data}))},handleGroup:function(){var e=this;this.templateList=this.templateAll.filter((function(t){if(t.groupId===e.projectTemplate.groupId)return t}))},getTemplateList:function(){var e=this;Object(T["f"])({}).then((function(t){e.templateAll=t.data.list}))},handleTableSwitch:function(e){var t=this;b(e).then((function(e){t.getList()}))},handleDetail:function(e){var t=this;this.dialogFormVisible=!0,this.projectTemplate={projectId:this.projectId,isGenerate:1,isOverride:1},e&&f(e).then((function(e){t.templateList=t.templateAll.filter((function(t){if(t.groupId===e.data.groupId)return t})),t.projectTemplate=e.data,t.projectTemplate.projectId=t.projectId}))},handleFormSubmit:function(e){var t=this;this.$refs[e].validate((function(e){e&&(t.projectTemplate.id?b(t.projectTemplate).then((function(e){t.dialogFormVisible=!1,t.getList()}),(function(){var e="";t.templateList.map((function(a){a.id===t.projectTemplate.templateId&&(e=a.name)})),t.$message.error(e+" "+t.$t("common.message.exists"))})):h(t.projectTemplate).then((function(e){t.dialogFormVisible=!1,t.getList()}),(function(){var e="";t.templateList.map((function(a){a.id===t.projectTemplate.templateId&&(e=a.name)})),t.$message.error(e+" "+t.$t("common.message.exists"))})))}))},handleFormClose:function(e){this.dialogFormVisible=!1,this.$refs[e].resetFields()},handleDelete:function(e){var t=this;this.$confirm(this.$t("common.confirm.deleteOne"),this.$t("common.confirm.title"),{cancelButtonText:this.$t("common.confirm.cancel"),confirmButtonText:this.$t("common.confirm.confirm"),type:"warning"}).then((function(){j(e).then((function(){t.getList()}))}))}}},$=v,w=a("2877"),k=Object(w["a"])($,p,m,!1,null,"9d6ebc22",null),x=k.exports,I={name:"Home",components:{TableList:u["a"],TemplateList:x},data:function(){return{projectId:"",tableList:[],currentTable:{id:"",projectId:"",tableName:"",tableType:null,domainName:"",domainDesc:""},loading:!1,generateResult:""}},created:function(){var e=this;this.$route.params.id?(this.projectId=this.$route.params.id,this.getTableList()):Object(n["c"])().then((function(t){e.projectId=t.data,e.getTableList()}))},methods:{getTableList:function(){var e=this,t={projectId:this.projectId};o(t).then((function(t){e.tableList=t.data,e.tableList.map((function(e){e.name=e.tableName,e.table=1===e.tableType}))}))},handleTableClick:function(e){var t=this;c(this.projectId,e.tableName).then((function(a){a.data.tableName?t.currentTable=a.data:t.currentTable=e,t.tableList=t.tableList.map((function(e){return e.active=!1,e.tableName===t.currentTable.tableName&&(e.id=t.currentTable.id,e.domainName=t.currentTable.domainName,e.domainDesc=t.currentTable.domainDesc,e.active=t.currentTable.active=!0),e}))}))},handleGenerate:function(){var e=this;if(this.currentTable.tableName){var t=this.$refs.templateList.list;0!==t.length?(this.loading=!0,this.generateResult="",s(this.projectId,[this.currentTable]).then((function(t){setTimeout((function(){e.loading=!1,e.generateResult=t.data}),1500)}))):this.$message.warning(this.$t("projectGenerate.button.templateWarning"))}else this.$message.warning(this.$t("projectGenerate.button.tableWarning"))},handleDetail:function(){this.$refs.templateList.handleDetail()}}},y=I,O=Object(w["a"])(y,r,i,!1,null,"0d40cefb",null);t["default"]=O.exports},a6fb:function(e,t,a){"use strict";var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-row",[a("el-col",{attrs:{span:4}},[a("div",{staticClass:"tree-list",on:{contextmenu:function(e){e.preventDefault()},click:e.closeRightMenu}},[e.showSearch?a("el-input",{staticClass:"filter-input",model:{value:e.filterText,callback:function(t){e.filterText=t},expression:"filterText"}}):e._e(),a("el-card",{class:{"filter-tree":!0,"has-input":e.showSearch,"has-menu":e.showMenu},attrs:{shadow:"never"}},e._l(e.filterData,(function(t){return a("div",{key:t.name,class:t.active?"item active":"item",on:{click:function(a){return e.itemClick(t)},contextmenu:function(a){return e.rightClick(t,a)}}},[e.showIcon?a("svg-icon",{staticClass:"icon",attrs:{"icon-class":t.table?"table":"view"}}):e._e(),a("span",{staticClass:"name"},[e._v(e._s(t.name))])],1)})),0),a("div",{directives:[{name:"show",rawName:"v-show",value:!e.showMenu&&e.showRightMenu&&e.clickedItem.rightMenu,expression:"!showMenu && showRightMenu && clickedItem.rightMenu"}],staticClass:"right-menu",style:{left:e.left+"px",top:e.top+"px"},on:{click:function(e){e.stopPropagation()}}},[a("el-button-group",[a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-plus"},on:{click:e.addItem}},[e._v(" "+e._s(e.$t("components.treeList.addItem"))+" ")]),a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-edit"},on:{click:e.updateItem}},[e._v(" "+e._s(e.$t("components.treeList.updateItem"))+" ")]),a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-delete"},on:{click:e.deleteItem}},[e._v(" "+e._s(e.$t("components.treeList.deleteItem"))+" ")])],1)],1),a("div",{directives:[{name:"show",rawName:"v-show",value:e.showMenu,expression:"showMenu"}],staticClass:"bottom-menu"},[a("el-button-group",[a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-plus"},on:{click:e.addItem}}),a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-edit"},on:{click:e.updateItem}}),a("el-button",{staticClass:"menu",attrs:{size:"mini",icon:"el-icon-delete"},on:{click:e.deleteItem}})],1)],1)],1)]),a("el-col",{attrs:{span:20}},[a("div",{staticClass:"tree-right"},[e._t("default")],2)])],1)},i=[],n=(a("4de4"),a("c975"),a("b0c0"),{name:"TreeList",props:{data:{type:Array,required:!0},showSearch:{type:Boolean,required:!1,default:!0},showMenu:{type:Boolean,required:!1,default:!1},showIcon:{type:Boolean,required:!1,default:!1}},data:function(){return{filterText:"",filterData:[],showRightMenu:!1,top:0,left:0,clickedItem:""}},watch:{data:function(){this.filterTable()},filterText:function(e){this.filterTable(e)}},methods:{filterTable:function(){var e=this;this.filterData=this.data.filter((function(t){return!e.filterText||(e.filterText?t.name.toLowerCase().indexOf(e.filterText.toLowerCase())>-1:void 0)}))},itemClick:function(e){this.showMenu||(this.closeRightMenu(),this.clickedItem=e,this.$emit("itemClick",e))},rightClick:function(e,t){this.showMenu||(this.top=t.clientY,this.left=t.clientX,this.clickedItem=e,this.showRightMenu=!0,this.$emit("rightClick",e))},closeRightMenu:function(){this.showRightMenu=!1},addItem:function(){this.$emit("addItem")},updateItem:function(){this.$emit("updateItem",this.clickedItem)},deleteItem:function(){this.$emit("deleteItem",this.clickedItem)}}}),l=n,o=(a("1307"),a("2877")),c=Object(o["a"])(l,r,i,!1,null,null,null);t["a"]=c.exports},c621:function(e,t,a){"use strict";a.d(t,"f",(function(){return i})),a.d(t,"c",(function(){return n})),a.d(t,"d",(function(){return l})),a.d(t,"g",(function(){return o})),a.d(t,"a",(function(){return c})),a.d(t,"e",(function(){return s})),a.d(t,"h",(function(){return u})),a.d(t,"b",(function(){return p}));var r=a("b775");function i(e){return Object(r["a"])({url:"/template/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/template/export",method:"get",responseType:"blob",params:e})}function l(e){return Object(r["a"])({url:"/template/dump",method:"post",data:e})}function o(e){return Object(r["a"])({url:"/template/select",method:"get",params:{id:e}})}function c(e){return Object(r["a"])({url:"/template/check",method:"post",data:e})}function s(e){return Object(r["a"])({url:"/template/insert",method:"post",data:e})}function u(e){return Object(r["a"])({url:"/template/update",method:"put",data:e})}function p(e){return Object(r["a"])({url:"/template/delete",method:"delete",params:{ids:e}})}},d9f8:function(e,t,a){"use strict";a.d(t,"d",(function(){return i})),a.d(t,"e",(function(){return n})),a.d(t,"a",(function(){return l})),a.d(t,"c",(function(){return o})),a.d(t,"f",(function(){return c})),a.d(t,"b",(function(){return s}));var r=a("b775");function i(e){return Object(r["a"])({url:"/templateGroup/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/templateGroup/select",method:"get",params:{id:e}})}function l(e){return Object(r["a"])({url:"/templateGroup/check",method:"post",data:e})}function o(e){return Object(r["a"])({url:"/templateGroup/insert",method:"post",data:e})}function c(e){return Object(r["a"])({url:"/templateGroup/update",method:"put",data:e})}function s(e){return Object(r["a"])({url:"/templateGroup/delete",method:"delete",params:{ids:e}})}}}]);