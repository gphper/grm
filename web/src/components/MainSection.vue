<template>
    <el-container>
        <el-main>
            <el-tabs
            v-model="activeIndex"
            closable
            type="card"
            @tab-click="handleClick"
            @tab-remove="removeTag"
            >
            <el-tab-pane
                v-for="item in tagData"
                :key="item.name"
                :label="item.title"
                :name="item.name"
            >
               {{ item.content }} 
            </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>

<script>
import { computed, watch } from "@vue/runtime-core";
import { useStore } from "vuex";
export default{

    name:"MainSection",
    setup(props) {

        const store = useStore()
        const tagData = computed(() => store.state.tagsList);
        const activeIndex = computed(() => store.state.activeTag);

        const handleClick = (tab) => {
            console.log('click '+ tab)
        };

        const removeTag = (index) => {
            store.commit("delTagsItem",index)
        };

        watch(
        () => props.activeTab,
        (newValue) => {
            activeIndex.value = newValue
        }
        );

        return { handleClick, removeTag, activeIndex,tagData };
    }
}
</script>

<style scoped>
.el-header{
    margin-top: 10px;
}
</style>