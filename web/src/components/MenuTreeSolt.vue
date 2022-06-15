<template>
    <div>
        <i v-if="data.children && data.children.length > 0" class="iconfont icon-folder-close"></i>
        <i v-else class="iconfont icon-key"></i>
        <span>&nbsp;{{ node.label }}</span>
        <span v-if="data.children && data.children.length > 0" style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="append(data)" type="success">添加</el-button>
            <el-button text @click.stop="remove(node,data)" type="danger">删除</el-button>
        </span>
        <span v-else style="margin-left:20%;width: 15px;">
            <el-button text @click.stop="remove(node,data)" type="danger">删除</el-button>
            <el-button text @click.stop="detail(node.label)" type="primary">查看</el-button>
        </span>
        <DataForm :visible="visible"></DataForm>
    </div>
</template>

<script>
import { useStore } from 'vuex'
import { ShowList, ShowString } from '@/utils/show.js'
import CryptoJS from "crypto-js";

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

        const append = () => {
            // const newChild = { id: 10, label: 'testtest', children: [] }
            // if (!data.children) {
            //     data.children = []
            // }
            // data.children.push(newChild)
            store.commit("switchDataForm");
        }

        const remove = (node,data) => {
            const parent = node.parent
            const children = parent.data.children || parent.data
            const index = children.findIndex((d) => d.id === data.id)
            children.splice(index, 1)
        }

        const detail = (key)=>{
            let id = CryptoJS.MD5(key).toString();
            store.commit("setTagsItem", {
                title: key,
                name: key,
                id: id
            });
            store.commit("setCurrentTag", key);

            //todo 根据key获取类型
            let types = "";    
            if(key == "string_key"){
                types = "string";
            }else{
                types = "list";
            }


            switch(types){
                case "string":
                    ShowString(key,id)
                    break;
                case "list":
                    ShowList(key,id)
                    break;    
            }
        }

        return{
            append,
            remove,
            detail
        }
    },
}
</script>