import { createStore } from 'vuex'

export default createStore({
  state: {
      tagsList: [],
      connList: [],
      activeTag: '',
      showDataForm: false,
      loadReq:false,
      sk:'',
      db:0,
  },
  mutations: {
      setCurrentTag(state,tag){
        state.activeTag = tag;
      },
      switchDataForm(state,obj){
        state.showDataForm = obj.show
        state.sk = obj.sk
        state.db = obj.db
      },
      delTagsItem(state,indexx) {

        let nextTag = {};
        let newTag = [];
        state.tagsList.forEach(function (item, index) {
            if (item.name == indexx) {
              //定位到指定值
              nextTag = state.tagsList[index - 1] || state.tagsList[index + 1];
            } else {
              newTag.push(item);
            }
        });

        if (newTag.length > 0) {
            state.activeTag = nextTag.name;
        }

        state.tagsList = newTag;
      },

      setTagsItem(state, data) {
          let exit = false;
          //先判断存不存在菜单
          state.tagsList.forEach(function (one) {
            if (one.name == data.name) {
              //定位到指定值
              exit = true;
            }
          });
    
          if (!exit) {
            state.tagsList.push(data)
          }
      },
      clearTags(state) {
          state.tagsList = []
      },
      closeTagsOther(state, data) {
          state.tagsList = data;
      },
      setLoadReq(state,load){
          state.loadReq = load
      },
      closeCurrentTag(state, data) {
          for (let i = 0, len = state.tagsList.length; i < len; i++) {
              const item = state.tagsList[i];
              if (item.path === data.$route.fullPath) {
                  if (i < len - 1) {
                      data
                          .$router
                          .push(state.tagsList[i + 1].path);
                  } else if (i > 0) {
                      data
                          .$router
                          .push(state.tagsList[i - 1].path);
                  } else {
                      data
                          .$router
                          .push("/");
                  }
                  state
                      .tagsList
                      .splice(i, 1);
                  break;
              }
          }
      },
      // 侧边栏折叠
      handleCollapse(state, data) {
          state.collapse = data;
      },
      addConn(state,data){
          state.connList.push(data)
      }
  },
  actions: {},
  modules: {}
})