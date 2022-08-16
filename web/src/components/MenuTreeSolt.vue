<template>
    <div>
        <span :title="node.label" class="showname">
        <i v-if="data.children && data.children.length > 0" class="iconfont icon-folder-close icon-style"></i>
        <i v-else class="iconfont icon-key icon-style"></i>
        &nbsp;{{ node.label }}
        </span>
        <span class="tree-node" v-if="data.children && data.children.length > 0" style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="append(data.sk,data.db,false,data.id)" type="success">添加</el-button>
            <el-button text @click.stop="remove(data.id,data.sk,data.db,1)" type="danger">删除</el-button>
        </span>
        <span class="tree-node" v-else style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="remove(data.id,data.sk,data.db,0)" type="danger">删除</el-button>
            <el-button text @click.stop="detail(node.label,data.id,data.sk,data.db)" type="primary">查看</el-button>
        </span>
    </div>
</template>

<script>
import { useStore } from 'vuex'
import { ShowHash, ShowList, ShowSet, ShowString,ShowZset,ShowStream } from '@/utils/show.js'
import { getKeyType,delKeys } from "@/api/index.js"
import CryptoJS from "crypto-js";
import { ElMessage } from 'element-plus';

export default {
    name:"MenuTreeSolt",
    props:{
        data:{
            type:Object,
        },
        node:{
            type:Object
        }
    },
    setup() {
        
        const store = useStore()

        const append = (sk,db,root,pre) => {
            store.commit("switchDataForm",{sk:sk,db:db,show:true,root:root,pre:pre});
        }

        const remove = (key,sk,db,single) => {
            delKeys({
                id:key,
                sk:sk,
                db:db,
                single:single
            }).then((res)=>{
                if(res.data.status == 1){
                    //删除成功
                    ElMessage({
                        message: "删除成功,请手动刷新列表",
                        type: 'success',
                    })
                }
            })
            store.commit("delTagsItem",sk+db)
        }

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

        return{
            append,
            remove,
            detail
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