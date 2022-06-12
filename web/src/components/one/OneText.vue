<template>
<div>
    <el-input v-model="currentData" :autosize="{ minRows:6, maxRows:12 }" type="textarea" placeholder="数据展示区" disabled />
</div>
    
</template>

<script>
import { ref } from '@vue/reactivity'
import { watch } from '@vue/runtime-core'

export default {
    name:"OneText",
    props:{
        data:{
            type:String
        }
    },
    setup(props) {
        let currentData = ref(props.data)

        watch(
            () => props.data,
            (value) => {
                currentData.value = value
            }
        )

        return {
            currentData
        }
    },
}
</script>