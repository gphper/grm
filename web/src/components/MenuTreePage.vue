<template>
    <el-row>
        <el-col :span="16" :offset="3">
            <el-button style="width:100%;" type="info" size="small" @click="nextPage(props.index,props.match,props.next)">下一页</el-button>
        </el-col>
    </el-row>
</template>

<script>
import {dbKeysList} from "@/utils/tree.js"
export default {
    name:"MenuTreePage",
    props:{
        next:{
            type:Number
        },
        index:{
            type:String
        },
        match:{
            type:String
        }
    },
    setup(props) {

        const nextPage = (index,match,course)=>{
            console.log(course)
            dbKeysList(index,match,course)
        }

        return{
            props,
            nextPage
        }
    },
}
</script>