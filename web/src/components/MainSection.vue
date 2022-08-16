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
                :label="item.title.substr(0,10)"
                :name="item.name"
            >
               <span v-if="item.content">{{ item.content }}</span>
               <span v-if="item.id"><div :id="item.id"></div></span>
            </el-tab-pane>
            </el-tabs>
        </el-main>
    </el-container>
</template>

<script>
import { computed } from "@vue/runtime-core";
import { useStore } from "vuex";
export default{

    name:"MainSection",
    setup() {

        const store = useStore()
        const tagData = computed(() => store.state.tagsList);
        const activeIndex = computed(() => store.state.activeTag);

        const handleClick = (tab) => {
            console.log('click '+ tab)
        };

        const removeTag = (index) => {
            store.commit("delTagsItem",index)
        };

        return { handleClick, removeTag, activeIndex,tagData };
    }
}
</script>

<style scoped>
.el-header{
    margin-top: 10px;
}
</style>