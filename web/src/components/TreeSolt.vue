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
    </div>
</template>

<script>
import { useStore } from 'vuex'



export default {
    name:"TreeSolt",
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

        const append = (data) => {
            const newChild = { id: 10, label: 'testtest', children: [] }
            if (!data.children) {
                data.children = []
            }
            data.children.push(newChild)
        }

        const remove = (node,data) => {
            const parent = node.parent
            const children = parent.data.children || parent.data
            const index = children.findIndex((d) => d.id === data.id)
            children.splice(index, 1)
        }

        const detail = (name)=>{
            store.commit("setTagsItem", {
                title: name,
                name: name,
                content: name,
            });
            store.commit("setCurrentTag", name);
        }

        return{
            append,
            remove,
            detail
        }
    },
}
</script>