import {Stage1v1View} from "./stage1v1/Stage1v1View";
//////////////
///////////////////

var routes = [
    {
        path: '/', name: 'home',
        components: {default: Stage1v1View}
    }
];

// import Vue = require('vue');
// import VueRouter = require('vue-router');
declare var VueRouter;
declare var Vue;
var router = new VueRouter({
    routes // short for routes: routes
});

router.afterEach((to, from) => {
    // var toPath = to.path;
    // router.app.active = toPath.split("/")[1];
    // router.app.monitorModel = monitorModel;
});

var app = new Vue({
    router
}).$mount('#app');
