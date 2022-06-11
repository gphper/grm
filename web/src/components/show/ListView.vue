<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <el-row :gutter="20">
                    <el-col :span="10">
                        <el-input  disabled :value="props.xkey">
                            <template #prepend>LIST:</template>
                        </el-input>
                    </el-col>
                    <el-col :span="10" :offset="2">
                        <el-button text bg>重命名</el-button>
                        <el-button text bg>TTL</el-button>
                        <el-button type="danger" text bg>删除</el-button>
                        <el-button type="warning" text bg>重载数据</el-button>
                    </el-col>
                </el-row>
            </template>

            <!-- 列表 -->
            <el-row>
                <el-col :span="20">
                    <el-table :data="filterTableData" height="250" border style="width: 100%" >
                        <el-table-column prop="num" label="Num" width="180" />
                        <el-table-column prop="value" label="Value"/>
                    </el-table>
                </el-col>
                <el-col :offset="1" :span="2">
                    <el-row class="data-button"><el-button plain type="primary">插入行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="danger">删除行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="success">重新载入</el-button></el-row>
                    <el-row class="data-button">
                        <el-input v-model="search" placeholder="值搜索" />
                    </el-row>
                </el-col>
            </el-row>


            <OneDetail :data="data"></OneDetail>
        </el-card>
    </div>
</template>

<style scoped>
.data-button{
    margin-top: 10px;
}
</style>

<script>
import OneDetail from "@/components/one/OneDetail.vue"
import { ref } from '@vue/reactivity';
import { computed } from '@vue/runtime-core';

export default {
    name: "ListView",
    props: {
        xkey: {
            type: String
        },
    },
    components:{
        OneDetail
    },
    setup(props) {
        //TODO 获取data值
        let data = `{"username":"gphper","nickname":"GPHER"}`;
        
        let search = ref('');

        let tableData = [
            {
                num: 1,
                value: 'Tom11',
            },
            {
                num: 2,
                value: 'Tom12',
            },
            {
                num: 3,
                value: 'Tom13'
            },
            {
                num: 4,
                value: 'Tom14'
            },
            {
                num: 5,
                value: 'Tom21'
            },
            {
                num: 6,
                value: 'Tom16'
            },
            {
                num: 7,
                value: 'Tom36'
            },
            {
                num: 8,
                value: 'Tom46'
            },
        ];

        let filterTableData = computed(()=>tableData.filter(
            (data) =>
            !search.value ||
            data.value.toLowerCase().includes(search.value.toLowerCase())
        ));

        return {
            props,
            data,
            tableData,
            search,
            filterTableData
        };
    },
}
</script>