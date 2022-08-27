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
                       
                        <el-tooltip
                            class="box-item"
                            effect="dark"
                            :content="ttl"
                            placement="top-start"
                        >
                            <el-button text bg @click="ttlVisible = true">TTL:<span style="width: 22px;overflow: hidden;">{{ttl}}</span></el-button>
                        </el-tooltip>
                        &nbsp;&nbsp;

                        <el-button type="danger" text bg @click="delKey(props.sk+props.db)">删除</el-button>
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
                        <el-table-column prop="score" label="Score" />
                        <el-table-column prop="value" label="Value"/>
                    </el-table>
                </el-col>
                <el-col :offset="1" :span="2">
                    <el-row class="data-button"><el-button plain type="primary" @click="addItemFormVisible = true">插入一行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="danger" @click="delItem">删除该行</el-button></el-row>
                    <el-row class="data-button"><el-button plain type="success" @click="reload">重新载入</el-button></el-row>
                    <el-row class="data-button">
                        <el-input v-model="search" placeholder="列表搜索" />
                    </el-row>
                    <el-row class="data-button">
                        <div class="page">
                            页:{{page}} 共{{total}}
                        </div>
                    </el-row>
                    <el-row class="data-button">
                        <el-col :span="12">
                            <el-button type="success" size="small" plain circle @click="prePage"><i class="iconfont icon-shangyiye"></i></el-button>
                        </el-col>
                        <el-col :span="12">
                            <el-button type="success" size="small" plain circle @click="nextPage"><i class="iconfont icon-xiayiye"></i></el-button>
                        </el-col>
                    </el-row>
                </el-col>
            </el-row>

            <OneDetail :data="data"></OneDetail>
            <TtlForm @close="close" @reload="reload" :visiable="ttlVisible" :sk="props.sk" :db="props.db" :keyx="props.xkey"></TtlForm>
        </el-card>

        <!-- 新增行 -->
        <el-dialog v-model="addItemFormVisible" @close="onReset" title="新增一行">
            <el-form
                :rules="addItemRules"
                ref="addItemFormRef"
                :model="addItemForm"
                label-width="120px"
                status-icon
            >

                <el-form-item label="分数" prop="score">
                    <el-input type="number" v-model="addItemForm.score" placeholder="分数" />
                </el-form-item>

                <el-form-item label="值" prop="item">
                    <el-input type="textarea" v-model="addItemForm.item" placeholder="新增值" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="addItemSubmit ">确定</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
        
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
import OneDetail from "@/components/one/OneDetail.vue";
import TtlForm from "@/components/one/TtlForm.vue";
import { reactive, ref } from '@vue/reactivity';
import store from '@/store/index.js';
import { delKeys } from "@/api/index.js";
import { computed, onMounted } from '@vue/runtime-core';
import { ElMessage } from 'element-plus';
import { showZset,delZset,addZset } from "@/api/zset.js"

export default {
    name: "ZsetView",
    props: {
        xkey: {
            type: String
        },
        sk:{
            type:String
        },
        db:{
            type:Number
        },
        node:{
            type:Object
        }
    },
    components:{
        OneDetail,
        TtlForm
    },
    setup(props,{emit}) {
        
        let limit = 1000;
        let ttl = ref(0);
        let page = ref(1);
        let total = ref(0);
        let data = ref('');
        let search = ref('');
        let tableData = ref([]);
        let ttlVisible = ref(false);
        
        //添加item
        let addItemFormVisible = ref(false);
        let addItemFormRef = ref(null)
        let addItemForm = reactive({
            id:props.xkey,
            sk:props.sk,
            db:props.db,
            score:0,
            item:"",
        })
        let addItemRules = {
            item:[
                {required:true,message: '请输入新增值', trigger: 'blur' }
            ],
        }
        

        let filterTableData = computed(()=>tableData.value.filter(
            (data) =>
            !search.value ||
            data.value.toLowerCase().includes(search.value.toLowerCase())
        ));

        let singleTableRef = ref(null)

        const rowClick = (row)=>{
            singleTableRef.value.setCurrentRow(row);
            data.value = row.value;
        }

        const close = ()=>{
            ttlVisible.value = false
        }

        const delKey = (key)=>{
            delKeys({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
            }).then((res)=>{
                if(res.data.status == 1){
                    //删除成功
                    ElMessage({
                        message: "删除成功",
                        type: 'success',
                    })
                    emit("del",props.node,props.node.data)
                    store.commit("delTagsItem",key)
                }
            })

            store.commit("delTagsItem",key)
        }

        const reload = ()=>{
            showZset({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                page:page.value,
                limit:limit
            }).then((res)=>{
                tableData.value = res.data.data
                ttl.value = res.data.ttl
                page.value = res.data.page
                total.value = res.data.total
            })
        }

        const prePage = ()=>{
            if (page.value > 1){
                page.value -= 1
            }
            showZset({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                page:page.value,
                limit:limit
            }).then((res)=>{
                tableData.value = res.data.data
                ttl.value = res.data.ttl
                page.value = res.data.page
                total.value = res.data.total
            })
        }

        const nextPage = ()=>{
            if (page.value < total.value){
                page.value += 1
            }
            showZset({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                page:page.value,
                limit:limit
            }).then((res)=>{
                tableData.value = res.data.data
                ttl.value = res.data.ttl
                page.value = res.data.page
                total.value = res.data.total
            })
        }

        const delItem = ()=>{
            delZset({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                item:data.value
            }).then((res)=>{
                if(res.data.count > 0){
                    ElMessage({
                        message: "删除成功",
                        type: 'success',
                    });
                    reload();
                }
            })
        }

        const addItemSubmit = ()=>{
            addZset(addItemForm).then((res)=>{
                if(res.data.count > 0){
                    ElMessage({
                        message: "添加成功",
                        type: 'success',
                    });
                    addItemFormVisible.value = false;
                    addItemFormRef.value.resetFields();
                    reload();
                }
            })
        }

        onMounted(reload)

        return {
            ttl,
            data,
            page,
            total,
            props,
            search,
            tableData,
            ttlVisible,
            singleTableRef,
            filterTableData,
            addItemFormVisible,
            addItemFormRef,
            addItemRules,
            addItemForm,
            close,
            reload,
            delKey,
            delItem,
            prePage,
            nextPage,
            rowClick,
            addItemSubmit
        };
    },
}
</script>