<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <el-row :gutter="20">
                    <el-col :span="12">
                        <el-input  disabled :value="props.xkey">
                            <template #prepend>STRING:</template>
                        </el-input>
                    </el-col>
                    <el-col :span="8" :offset="2">
                        <el-button text bg>重命名</el-button>
                        <el-button text bg>TTL</el-button>
                        <el-button type="danger" text bg>删除</el-button>
                        <el-button type="warning" text bg>重载数据</el-button>
                    </el-col>
                </el-row>
            </template>
            
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
            <component :is="currentView" :data="data"></component>
            
        </el-card>
    </div>
</template>

<style scoped>
.toolbox{
    padding-left: 10px;
    padding-right: 20px;
    padding-bottom: 10px;
    width: 100%;
}
</style>

<script>
import { ref } from '@vue/reactivity'
import OneText from '@/components/OneText.vue'
import OneJson from '@/components/OneJson.vue'
import OneBase64 from '@/components/OneBase64.vue'
import { computed } from '@vue/runtime-core'
import { Buffer } from 'buffer'

export default {
    name: "StringView",
    components: {
        OneText,
        OneJson,
        OneBase64
    },
    props: {
        xkey: {
            type: String
        },
    },
    setup(props) {
        //TODO 获取data值
        let data = ref("");
        data = "world";
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
            get:()=>Buffer.byteLength(data)
        })

        return {
            props,
            data,
            viewOption,
            currentView,
            size
        };
    },
}
</script>