<template>
    <div class="common-layout">
    <el-container>
      <el-header>
        <el-row style="margin-top:15px;">
            <el-col :span="6">
                <el-switch
                    v-model="showswitch"
                    class="mb-2"
                    style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
                    active-text="表格"
                    inactive-text="图表"
                />
            </el-col>

            <el-col :span="4" :offset="6">
                <el-form-item label="是否自动更新">
                    <el-switch
                        @change="autoChange"
                        v-model="autoupdate"
                        inline-prompt
                        active-text="是"
                        inactive-text="否"
                    />
                </el-form-item>
            </el-col>

            <el-col :span="5">
                <el-form-item label="更新间隔">
                    <el-select @change="intervalChange" v-model="interval" placeholder="间隔">
                        <el-option label="5s" value="5" />
                        <el-option label="10s" value="10" />
                        <el-option label="30s" value="30" />
                        <el-option label="1min" value="60" />
                        <el-option label="5min" value="300" />
                    </el-select>
                </el-form-item>
            </el-col>

            <el-col :span="2" :offset="1">
                <el-button type="success" title="重载" @click="getInfo" circle><i class="iconfont icon-zhongqi"></i></el-button>
            </el-col>

        </el-row>
      </el-header>
      <el-main>
            <div v-show="!showswitch">
                <el-row>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart1"></div>
                    </el-col>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart2"></div>
                    </el-col>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart3"></div>
                    </el-col>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart4"></div>
                    </el-col>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart5"></div>
                    </el-col>
                    <el-col :span="8">
                        <div style="width:350px;height:350px" id="chart6"></div>
                    </el-col>
                </el-row>
            
            </div>
            <div v-show="showswitch">
                <el-tabs tab-position="left" style="height: 100%" class="demo-tabs">
                    <el-tab-pane label="CPU">
                        <el-table :data="cpuData" style="width: 100%">
                            <el-table-column prop="name" label="NAME" width="250" />
                            <el-table-column prop="value" label="VALUE" width="250" />
                        </el-table>
                    </el-tab-pane>
                    <el-tab-pane label="Clients">
                        <el-table :data="clientData" style="width: 100%">
                            <el-table-column prop="name" label="NAME" width="250" />
                            <el-table-column prop="value" label="VALUE" width="250" />
                        </el-table>
                    </el-tab-pane>
                    <el-tab-pane label="Cluster">
                        <el-table :data="clusterData" style="width: 100%">
                            <el-table-column prop="name" label="NAME" width="250" />
                            <el-table-column prop="value" label="VALUE" width="250" />
                        </el-table>
                    </el-tab-pane>
                    <el-tab-pane label="Keyspace">
                        <el-table :data="keyspaceData" style="width: 100%">
                            <el-table-column prop="db" label="DB" width="250" />
                            <el-table-column prop="avg_ttl" label="AVG" width="250" />
                            <el-table-column prop="expires" label="EXPIRES" width="250" />
                            <el-table-column prop="keys" label="KEYS" width="250" />
                        </el-table>
                    </el-tab-pane>
                    <el-tab-pane label="Memory">
                        <el-scrollbar height="500px">
                            <el-table :data="memoryData" style="width: 100%">
                                <el-table-column prop="name" label="NAME" width="250" />
                                <el-table-column prop="value" label="VALUE" width="250" />
                            </el-table>
                        </el-scrollbar>
                    </el-tab-pane>
                    <el-tab-pane label="Persistence">
                        <el-scrollbar height="500px">
                            <el-table :data="persData" style="width: 100%">
                                <el-table-column prop="name" label="NAME" width="250" />
                                <el-table-column prop="value" label="VALUE" width="250" />
                            </el-table>
                        </el-scrollbar>
                    </el-tab-pane>
                    <el-tab-pane label="Replication">
                        <el-table :data="repData" style="width: 100%">
                            <el-table-column prop="name" label="NAME" width="250" />
                            <el-table-column prop="value" label="VALUE" width="250" />
                        </el-table>
                    </el-tab-pane>
                    <el-tab-pane label="Server">
                        <el-scrollbar height="500px">
                            <el-table :data="serverData" style="width: 100%">
                                <el-table-column prop="name" label="NAME" width="250" />
                                <el-table-column prop="value" label="VALUE" width="250" />
                            </el-table>
                        </el-scrollbar>
                    </el-tab-pane>
                    <el-tab-pane label="Stats">
                        <el-scrollbar height="500px">
                            <el-table :data="statsData" style="width: 100%">
                                <el-table-column prop="name" label="NAME" width="250" />
                                <el-table-column prop="value" label="VALUE" width="250" />
                            </el-table>
                        </el-scrollbar>
                    </el-tab-pane>
                </el-tabs>
            </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { computed, onMounted, onUnmounted, ref, watch } from '@vue/runtime-core'
import { useRoute } from 'vue-router'
import { serInfo } from "@/api/index.js"
import * as echarts from "echarts";
import { useStore } from 'vuex';
export default {
    name:"MonitorView",
    props:{
        sk:{
            type:String
        }
    },
    setup() {

        const store = useStore()
        const route = useRoute()

        let timer = null
        let xtime = []
        //内存
        let memuse = []
        let memChart = null
        let memChartOption = []
        const memChartChange = (used_memory)=>{
            if(memuse.length < 100){
                memuse.push(used_memory)
            }else{
                memuse.shift()
                memuse.push(used_memory)
            }
            memChartOption.xAxis.data = xtime
            memChartOption.series[0].data = memuse
            memChart.setOption(memChartOption)
        }


        //进口流量
        let input_kbps = []
        let inputChart = null
        let inputChartOption = []
        const inputChartChange = (input_kbps_val)=>{
            if(input_kbps.length < 100){
                input_kbps.push(input_kbps_val)
            }else{
                input_kbps.shift()
                input_kbps.push(input_kbps_val)
            }
            inputChartOption.xAxis.data = xtime
            inputChartOption.series[0].data = input_kbps
            inputChart.setOption(inputChartOption)
        }

        //出口流量
        let output_kbps = []
        let outputChart = null
        let outputChartOption = []
        const outputChartChange = (output_kbps_val)=>{
            if(output_kbps.length < 100){
                output_kbps.push(output_kbps_val)
            }else{
                output_kbps.shift()
                output_kbps.push(output_kbps_val)
            }
            outputChartOption.xAxis.data = xtime
            outputChartOption.series[0].data = output_kbps
            outputChart.setOption(outputChartOption)
        }

        //客户端连接数
        let connected_clients = []
        let connChart = null
        let connChartOption = []
        const connChartChange = (connected_client)=>{
            if(connected_clients.length < 100){
                connected_clients.push(connected_client)
            }else{
                connected_clients.shift()
                connected_clients.push(connected_client)
            }
            connChartOption.xAxis.data = xtime
            connChartOption.series[0].data = connected_clients
            connChart.setOption(connChartOption)
        }

        //每秒处理命令数 instantaneous_ops_per_sec
        let ops_per_secs = []
        let opsChart = null
        let opsChartOption = []
        const opsChartChange = (ops_per_sec)=>{
            if(ops_per_secs.length < 100){
                ops_per_secs.push(ops_per_sec)
            }else{
                ops_per_secs.shift()
                ops_per_secs.push(ops_per_sec)
            }
            opsChartOption.xAxis.data = xtime
            opsChartOption.series[0].data = ops_per_secs
            opsChart.setOption(opsChartOption)
        }

        //总key值数
        let keys_count = []
        let keysChart = null
        let keysChartOption = []
        const keysChartChange = (key_count)=>{
            if(keys_count.length < 100){
                keys_count.push(key_count)
            }else{
                keys_count.shift()
                keys_count.push(key_count)
            }
            keysChartOption.xAxis.data = xtime
            keysChartOption.series[0].data = keys_count
            keysChart.setOption(keysChartOption)
        }

        //cup
        let cupInfo = ref({})
        let cpuData = computed(()=>{
            let cpuArray = []
            for(let cpuObj in cupInfo.value){
                cpuArray.push({
                    "name":cpuObj,
                    "value":cupInfo.value[cpuObj]
                })
            }
            return cpuArray
        })

        //clients
        let clientInfo = ref({})
        let clientData = computed(()=>{
            let clientArray = []
            for(let clientObj in clientInfo.value){
                clientArray.push({
                    "name":clientObj,
                    "value":clientInfo.value[clientObj]
                })
            }
            return clientArray
        })

        //Cluster
        let clusterInfo = ref({})
        let clusterData = computed(()=>{
            let clusterArray = []
            for(let clusterObj in clusterInfo.value){
                clusterArray.push({
                    "name":clusterObj,
                    "value":clusterInfo.value[clusterObj]
                })
            }
            return clusterArray
        })

        //Memory
        let memoryInfo = ref({})
        let memoryData = computed(()=>{
            let memoryArray = []
            for(let memoryObj in memoryInfo.value){
                memoryArray.push({
                    "name":memoryObj,
                    "value":memoryInfo.value[memoryObj]
                })
            }
            return memoryArray
        })

        //Persistence
        let persInfo = ref({})
        let persData = computed(()=>{
            let persArray = []
            for(let persObj in persInfo.value){
                persArray.push({
                    "name":persObj,
                    "value":persInfo.value[persObj]
                })
            }
            return persArray
        })

        //Replication
        let repInfo = ref({})
        let repData = computed(()=>{
            let repArray = []
            for(let repObj in repInfo.value){
                repArray.push({
                    "name":repObj,
                    "value":repInfo.value[repObj]
                })
            }
            return repArray
        })

        //Server
        let serverInfo = ref({})
        let serverData = computed(()=>{
            let serverArray = []
            for(let serverObj in serverInfo.value){
                serverArray.push({
                    "name":serverObj,
                    "value":serverInfo.value[serverObj]
                })
            }
            return serverArray
        })

        //Stats
        let statsInfo = ref({})
        let statsData = computed(()=>{
            let statsArray = []
            for(let statsObj in statsInfo.value){
                statsArray.push({
                    "name":statsObj,
                    "value":statsInfo.value[statsObj]
                })
            }
            return statsArray
        })

        //Keyspace
        let keyspaceInfo = ref({})
        let keyspaceData = computed(()=>{
            let keyspaceArray = []
            for(let keyspaceObj in keyspaceInfo.value){
                keyspaceArray.push({
                    "db":keyspaceObj,
                    "avg_ttl":keyspaceInfo.value[keyspaceObj].avg_ttl,
                    "expires":keyspaceInfo.value[keyspaceObj].expires,
                    "keys":keyspaceInfo.value[keyspaceObj].keys,
                })
            }
            return keyspaceArray
        })



        const trans = (size) => {
            console.log(size)
            let kb = 1024
            let unit = ['B','KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
            let i = Math.floor(Math.log(size) / Math.log(kb))
            if(size > 0){
                return (size / Math.pow(kb, i)).toPrecision(3) + ' ' + unit[i];
            }else{
                return '0B'
            }
        }

        const getInfo = ()=>{
            serInfo({key:route.params.sk}).then((res)=>{
                if(xtime.length < 100){
                    xtime.push(res.data.time)
                }else{
                    xtime.shift()
                    xtime.push(res.data.time)
                }
                
                //更新内存
                memChartChange(res.data.info.Memory.used_memory)
                //更新进口流量
                inputChartChange(res.data.info.Stats.instantaneous_input_kbps)
                //更新出口流量
                outputChartChange(res.data.info.Stats.instantaneous_output_kbps)
                //更新连接数
                connChartChange(res.data.info.Clients.connected_clients)
                //每秒处理数
                opsChartChange(res.data.info.Stats.instantaneous_ops_per_sec)
                //key值总数
                let count = 0;
                for(let obj in res.data.info.Keyspace){
                    count += res.data.info.Keyspace[obj].keys
                }
                keysChartChange(count)

                cupInfo.value = res.data.info.CPU
                clientInfo.value = res.data.info.Clients
                clusterInfo.value = res.data.info.Cluster
                memoryInfo.value = res.data.info.Memory
                persInfo.value = res.data.info.Persistence
                repInfo.value = res.data.info.Replication
                serverInfo.value = res.data.info.Server
                statsInfo.value = res.data.info.Stats
                keyspaceInfo.value = res.data.info.Keyspace
            })
        }


        watch(
            () => store.state.activeTag,
            (value) => {
                if(value != "grmmonitor"){
                    clearInterval(timer)
                    timer = null
                }else{
                    timer = setInterval(()=>{
                        getInfo()
                    },interval.value * 1000)
                }
            }
        )


        const intervalChange = (val)=>{
            if(autoupdate.value){
                clearInterval(timer)
                timer = setInterval(()=>{
                    getInfo()
                },val * 1000)
            }
        }

        onUnmounted(()=>{
            clearInterval(timer)
            timer = null
        })

        const autoChange = (val) => {
            if(val == false){
                clearInterval(timer)
                timer = null
            }else{
                timer = setInterval(()=>{
                    getInfo()
                },interval.value * 1000)
            }
        }

       let showswitch =  ref(false)
       let autoupdate = ref(true)
       let interval = ref(5)


       if(autoupdate.value){
            timer = setInterval(()=>{
                getInfo()
            },interval.value * 1000)
       }

       onMounted(() => {
            xtime = []
            memChart = echarts.init(document.getElementById('chart1'));
            memChartOption = {
                title: {
                    text: 'Redis内存使用'
                },
                grid:{
                    left:50,// 调整这个属性
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ trans(params[0].data);
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                    axisLabel:{
                        formatter:(value) => {
                            return trans(value)
                        }
                    }
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            memChart.setOption(memChartOption);


            inputChart = echarts.init(document.getElementById('chart2'));
            inputChartOption = {
                title: {
                    text: 'Redis进口流量'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ params[0].data +" kbps";
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            inputChart.setOption(inputChartOption);

            outputChart = echarts.init(document.getElementById('chart3'));
            outputChartOption = {
                title: {
                    text: 'Redis出口流量'
                },
                tooltip: {
                    trigger: 'axis',
                    formatter:(params)=>{
                        return params[0].axisValueLabel+":"+ params[0].data +" kbps";
                    },
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            outputChart.setOption(outputChartOption);

            connChart = echarts.init(document.getElementById('chart4'));
            connChartOption = {
                title: {
                    text: 'Redis连接数'
                },
                tooltip: {
                    trigger: 'axis',
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            connChart.setOption(connChartOption);


            opsChart = echarts.init(document.getElementById('chart5'));
            opsChartOption = {
                title: {
                    text: 'Redis每秒处理数'
                },
                tooltip: {
                    trigger: 'axis',
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            opsChart.setOption(opsChartOption);

            keysChart = echarts.init(document.getElementById('chart6'));
            keysChartOption = {
                title: {
                    text: 'Redis键值总数'
                },
                tooltip: {
                    trigger: 'axis',
                },
                xAxis: {
                    type: 'category',
                    data: []
                },
                yAxis: {
                    type: 'value',
                },
                series: [
                    {
                        data: [],
                        type: 'line',
                        smooth: true
                    }
                ]
            };
            keysChart.setOption(keysChartOption);
       })

       return {
            showswitch,
            interval,
            autoupdate,
            cpuData,
            clientData,
            clusterData,
            memoryData,
            statsData,
            serverData,
            repData,
            persData,
            keyspaceData,
            getInfo,
            autoChange,
            intervalChange
       }
    },
}
</script>

<style scoped>
.el-header{
    background-color:  #EBEDF0;
}
</style>