<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <el-row :gutter="20">
                    <el-col :span="16">
                        <el-input  disabled :value="props.xkey">
                            <template #prepend>LIST:</template>
                        </el-input>
                    </el-col>
                    <el-col :span="6" :offset="2">
                        <el-button text bg>重命名</el-button>
                        <el-button text bg>TTL</el-button>
                        <el-button type="danger" text bg>删除</el-button>
                    </el-col>
                </el-row>
            </template>

            <!-- 列表 -->
            <el-row class="list-panel">
                <el-col :span="20">
                    <el-table 
                    :data="filterTableData"  
                    height="250" 
                    highlight-current-row
                    border 
                    style="width: 100%"
                    ref="singleTableRef"
                    @row-click="rowClick"
                    >
                        <el-table-column prop="num" label="Num" width="180" />
                        <el-table-column prop="value" label="Value"/>
                    </el-table>
                </el-col>
                <el-col :offset="1" :span="2">
                    <el-row class="data-button"><el-button plain type="primary">插入一行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="danger">删除该行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="success">重新载入</el-button></el-row>
                    <el-row class="data-button">
                        <el-input v-model="search" placeholder="列表搜索" />
                    </el-row>
                    <el-row class="data-button">
                        <div class="page">
                            页:1 共10
                        </div>
                    </el-row>
                    <el-row class="data-button">
                        <el-col :span="12">
                            <el-button type="success" size="small" plain circle><i class="iconfont icon-shangyiye"></i></el-button>
                        </el-col>
                        <el-col :span="12">
                            <el-button type="success" size="small" plain circle><i class="iconfont icon-xiayiye"></i></el-button>
                        </el-col>
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

.list-panel{
    padding-bottom: 15px;
    margin-top: 15px;
    border-bottom: 1px solid #E4E7ED;
}

.page{
    font-size: 10px;
    color:#F0F2F5;
    background-color: #C0C4CC;
    padding: 8px;
    border-radius: 5px;
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
        let data = ref('');
        
        let search = ref('');

        let tableData = [
            {
                num: 1,
                value: 'Tom1133',
            },
            {
                num: 2,
                value: 'Tom12232',
            },
            {
                num: 3,
                value: 'Tom13444'
            },
            {
                num: 4,
                value: 'Tom1'
            },
            {
                num: 5,
                value: 'Tom21'
            },
            {
                num: 6,
                value: 'To3m16'
            },
            {
                num: 7,
                value: 'Tom3eeeeeeeee6'
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

        let singleTableRef = ref(null)

        const rowClick = (row)=>{
            singleTableRef.value.setCurrentRow(row);
            data.value = row.value;
        }

        return {
            props,
            data,
            tableData,
            search,
            filterTableData,
            singleTableRef,
            rowClick
        };
    },
}
</script>