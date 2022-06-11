<template>
    <div>
        <div class="toolbox">
            <el-row>
                <el-col :span="3">
                    <span class="ml-3 w-35 text-gray-600 inline-flex items-center">键值：{{size}} Byte</span>
                </el-col>
                
                <el-col :span="8" :offset="10">
                    <span class="ml-3 w-35 text-gray-600 inline-flex items-center">查看：</span>
                    <el-select v-model="currentView" class="m-2" placeholder="Select">
                        <el-option
                        v-for="item in viewOption"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        />
                    </el-select>
                </el-col>
            </el-row>

        </div>
        <component :is="currentView" :data="props.data"></component>
    </div>
</template>

<style scoped>
.toolbox{
    padding-left: 10px;
    padding-right: 20px;
    padding-bottom: 10px;
    padding-top: 10px;
    width: 100%;
}
</style>

<script>
import { ref } from '@vue/reactivity'
import { computed } from '@vue/runtime-core'
import { Buffer } from 'buffer'
import OneText from '@/components/one/OneText.vue'
import OneJson from '@/components/one/OneJson.vue'
import OneBase64 from '@/components/one/OneBase64.vue'

export default {
    name:"OneDetail",
    props:{
        data:{
            type:String
        },
    },
    components:{
        OneText,
        OneJson,
        OneBase64
    },
    setup(props) {
    
        let viewOption = ref([
            {
                value: "OneText",
                label: "Plain Text",
            },
            {
                value: "OneJson",
                label: "JSON",
            },
            {
                value: "OneBase64",
                label: "BASE64",
            }
        ]);
        let currentView = ref("OneText");

        let size = computed({
            get:()=>Buffer.byteLength(props.data)
        })

        return {
            props,
            viewOption,
            currentView,
            size
        };
    },
}
</script>