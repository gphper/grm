<template>

    <el-tree-v2
      :data="datas"
      node-key="id"
      default-expand-all
      :expand-on-click-node="false"
    >
      <template #default="{ node, data }">
        <span :title="node.label" class="showname">
        <i v-if="data.children && data.children.length > 0" class="iconfont icon-folder-close icon-style"></i>
        <i v-else class="iconfont icon-key icon-style"></i>
        &nbsp;{{ node.label }}
        </span>
        <span v-if="data.children && data.children.length > 0" style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="append(data)" type="success">添加</el-button>
            <el-button text @click.stop="remove(data.id,data.sk,data.db,1,node, data)" type="danger">删除</el-button>
        </span>
        <span v-else style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="remove(data.id,data.sk,data.db,0,node, data)" type="danger">删除</el-button>
            <el-button text @click.stop="detail(node.label,data.id,data.sk,data.db)" type="primary">查看</el-button>
        </span>
      </template>
    </el-tree-v2>
</template>

<script>
import { ref } from '@vue/reactivity'
import { getKeys,delKeys } from '@/api/index.js'
import { useStore } from 'vuex'
import { ShowHash, ShowList, ShowSet, ShowString,ShowZset,ShowStream } from '@/utils/show.js'
import { getKeyType } from "@/api/index.js"
import CryptoJS from "crypto-js";
import { ElMessage } from 'element-plus';

export default {
    name:"MenuTree",
    props:{
      index:{
        type:String
      }
    },
    setup(props) {

      let datas = ref([]);
      const store = useStore()

      const append = (data) => {
            let newChild = { id: 30, label: 'testtest', children: [] }
            if (!data.children) {
              data.children = []
            }
            data.children.push(newChild)
            datas.value = [...datas.value]
      };

      const remove = (key,sk,db,single,node, data)=>{
          delKeys({
                id:key,
                sk:sk,
                db:db,
                single:single
          }).then((res)=>{
              if(res.data.status == 1){
                  //删除成功
                  ElMessage({
                      message: "删除成功",
                      type: 'success',
                  })

                  let parent = node.parent
                  if(!parent){
                    parent = datas.value
                    let index = parent.findIndex((d) => d.id === data.id)
                    parent.splice(index, 1)
                  }else{
                    let children = parent.data.children || parent.data
                    let index = children.findIndex((d) => d.id === data.id)
                    children.splice(index, 1)
                  }
                  datas.value = [...datas.value]
              }
          })
          store.commit("delTagsItem",sk+db)
      };

      const detail = (key,idk,sk,db)=>{
            let unionid = sk+db
            let id = CryptoJS.MD5(unionid).toString();
            
            getKeyType({"id":idk,"sk":sk,"db":db}).then((res)=>{
                switch(res.data.types){
                    case "string":
                        ShowString(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    case "list":
                        ShowList(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    case "hash":
                        ShowHash(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    case "set":
                        ShowSet(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    case "zset":
                        ShowZset(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    case "stream":
                        ShowStream(idk,id,sk,db);
                        store.commit("setTagsItem", {
                            title: key,
                            name: unionid,
                            id: id
                        });
                        store.commit("setCurrentTag", unionid);
                        break;
                    default:
                        ElMessage({
                            message: "该键值不存在",
                            type: 'warning',
                        })
                        break;                 
                }
            })
        }


        getKeys({"index":props.index,"match":"*","cursor":0}).then((res) => {
          document.getElementById("menu#"+props.index).nextElementSibling.setAttribute("style","display:flex")
          datas.value = res.data.data
          if(res.data.cursor != 0){
              document.getElementById("page#"+props.index).style.display="block"
          }else{
              document.getElementById("page#"+props.index).style.display="none"
          }
          document.getElementById("dbnum#"+props.index).innerHTML = res.data.count;
        })


      // onMounted(()=>{
      //   getKeys({"index":props.index,"match":"*","cursor":0}).then((res) => {
      //     document.getElementById("menu#"+props.index).nextElementSibling.setAttribute("style","display:flex")
      //     datas.value = res.data.data
      //     if(res.data.cursor != 0){
      //         document.getElementById("page#"+props.index).style.display="block"
      //     }else{
      //         document.getElementById("page#"+props.index).style.display="none"
      //     }
      //     document.getElementById("dbnum#"+props.index).innerHTML = res.data.count;
      // });

      

      return {
        append,
        remove,
        detail,
        datas
      }
    },
}
</script>

<style scoped>
.tree-node{
    height: 38px;
}

.showname{
    width: 170px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    display: inline-block;
    vertical-align:middle;
}

</style>