<template>
<div>
    <el-input v-model="currentData" :autosize="{ minRows:6, maxRows:12 }" type="textarea" placeholder="Please input" disabled />
</div>
    
</template>

<script>
import { ref } from '@vue/reactivity'

export default {
    name:"OneText",
    props:{
        data:{
            type:String
        }
    },
    setup(props) {
        let currentData = ref(props.data)

        return {
            currentData
        }
    },
}
</script>