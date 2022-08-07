<template>
    <div>
        <i v-if="data.children && data.children.length > 0" class="iconfont icon-folder-close"></i>
        <i v-else class="iconfont icon-key"></i>
        <span>&nbsp;{{ node.label }}</span>
        <span v-if="data.children && data.children.length > 0" style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="append(data.sk,data.db,false,data.id)" type="success">添加</el-button>
            <el-button text @click.stop="remove(data.id)" type="danger">删除</el-button>
        </span>
        <span v-else style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="remove(data.id)" type="danger">删除</el-button>
            <el-button text @click.stop="detail(node.label,data.id,data.sk,data.db)" type="primary">查看</el-button>
        </span>
        <DataForm></DataForm>
    </div>
</template>

<script>
import { useStore } from 'vuex'
import DataForm from '@/components/index/DataForm.vue'
import { ShowHash, ShowList, ShowSet, ShowString,ShowZset,ShowStream } from '@/utils/show.js'
import { getKeyType } from "@/api/index.js"
import CryptoJS from "crypto-js";

export default {
    name:"MenuTreeSolt",
    components:{
        DataForm
    },
    props:{
        data:{
            type:Object,
        },
        node:{
            type:Object
        }
    },
    setup(props) {
        
        const store = useStore()

        const append = (sk,db,root,pre) => {
            store.commit("switchDataForm",{sk:sk,db:db,show:true,root:root,pre:pre});
        }

        const remove = (key) => {
            const parent = props.node.parent
            const children = parent.data.children || parent.data
            const index = children.findIndex((d) => d.id === key)
            children.splice(index, 1)
        }


        const detail = (key,idk,sk,db)=>{
            let unionid = sk+db
            let id = CryptoJS.MD5(unionid).toString();
            store.commit("setTagsItem", {
                title: key,
                name: unionid,
                id: id
            });

            getKeyType({"id":idk,"sk":sk,"db":db}).then((res)=>{
                switch(res.data.types){
                    case "string":
                        ShowString(idk,id,sk,db,remove)
                        break;
                    case "list":
                        ShowList(idk,id,sk,db,remove)
                        break;
                    case "hash":
                        ShowHash(idk,id,sk,db,remove)
                        break;
                    case "set":
                        ShowSet(idk,id,sk,db,remove)
                        break;
                    case "zset":
                        ShowZset(idk,id,sk,db,remove)
                        break; 
                    case "stream":
                        ShowStream(idk,id,sk,db,remove)
                        break;              
                }
            })
            store.commit("setCurrentTag", unionid);
        }

        return{
            append,
            remove,
            detail
        }
    },
}
</script>