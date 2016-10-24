/**
 * Created by toramisu on 2016/10/22.
 */
import {Navbar} from "./navbar/Navbar";
import {HomeView} from "./home/home";
//////////////
///////////////////

var routes = [
    {
        path: '/', name: 'home',
        components: {default: HomeView, Navbar: Navbar}
    },
    // {
    //     path: '/setting', name: 'setting',
    //     components: {default: SettingView, Navbar: Navbar},
    // }
];

declare var VueRouter;
declare var Vue;
var router = new VueRouter({
    routes // short for routes: routes
});

router.afterEach((to, from) => {
    var toPath = to.path;
    router.app.active = toPath.split("/")[1];
    // router.app.monitorModel = monitorModel;
});

var app = new Vue({
    router
}).$mount('#app');
