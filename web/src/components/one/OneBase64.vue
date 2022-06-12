<template>
    <div>
        <el-input v-model="currentData" :autosize="{ minRows:6, maxRows:12 }" type="textarea" placeholder="数据展示区" disabled />
    </div>
</template>

<script>
import { computed } from '@vue/runtime-core'
import { Buffer } from 'buffer'

export default {
    name:"OneBase64",
    props:{
        data:{
            type:String
        }
    },
    setup(props) {
        let currentData = computed({
            get:()=>{
                let tmp = Buffer.from(props.data,'base64').toString()
                return tmp;
            }
        })
        return {
            currentData
        }
    },
}
</script>