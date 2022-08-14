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
                            :content="ttl+''"
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
                        <el-col  :span="24" class="page">
                            共{{total}}个值
                        </el-col>
                    </el-row>
                    <el-row class="data-button">
                        <el-col :span="24">
                            <el-button v-if="cursor > 0" style="width:100%;" type="info" size="small" @click="nextPage">下一页</el-button>
                        </el-col>
                    </el-row>
                </el-col>
            </el-row>

            <OneDetail :data="data"></OneDetail>
            <TtlForm @close="close" @reload="reload" :visiable="ttlVisible" :sk="props.sk" :db="props.db" :keyx="props.xkey"></TtlForm>
        </el-card>

        <!-- 新增行 -->
        <el-dialog v-model="addItemFormVisible" title="新增一行">
            <el-form
                :rules="addItemRules"
                ref="addItemFormRef"
                :model="addItemForm"
                label-width="120px"
                status-icon
            >
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
    text-align: center;
}
</style>

<script>
import OneDetail from "@/components/one/OneDetail.vue";
import TtlForm from "@/components/one/TtlForm.vue";
import { reactive, ref } from '@vue/reactivity';
import store from '@/store/index.js';
import { delKeys } from "@/api/index.js";
import { onMounted, computed } from '@vue/runtime-core';
import { ElMessage } from 'element-plus';
import { showSet,delSet,addSet } from "@/api/set.js"

export default {
    name: "SetView",
    props: {
        xkey: {
            type: String
        },
        sk:{
            type:String
        },
        db:{
            type:Number
        }
    },
    components:{
        OneDetail,
        TtlForm
    },
    setup(props) {
        
        let ttl = ref(0);
        let total = ref(0);
        let cursor = ref(0);
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
                        message: "删除成功,请手动刷新列表更新",
                        type: 'success',
                    })
                }
            })

            store.commit("delTagsItem",key)
        }

        const reload = ()=>{
            showSet({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                cursor:0
            }).then((res)=>{
                tableData.value = res.data.data
                ttl.value = res.data.ttl
                total.value = res.data.total
                cursor.value = res.data.cursor
            })
        }

        const nextPage = ()=>{
            showSet({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
                cursor:cursor.value
            }).then((res)=>{
                tableData.value = res.data.data
                ttl.value = res.data.ttl
                total.value = res.data.total
                cursor.value = res.data.cursor
            })
        }

        const delItem = ()=>{
            delSet({
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
            addSet(addItemForm).then((res)=>{
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
            total,
            cursor,
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
            nextPage,
            rowClick,
            addItemSubmit
        };
    },
}
</script>