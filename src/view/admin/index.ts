/**
 * Created by toramisu on 2016/10/22.
 */
import {Navbar} from "./navbar/Navbar";
import {HomeView} from "./home/home";
import {PlayerView} from "./player/player";
import {RankView} from "./rank/rank";
//////////////
///////////////////

var routes = [
    {
        path: '/', name: 'home',
        components: {content: HomeView, Navbar: Navbar}
    },
    {
        path: '/player', name: 'player',
        components: {content: PlayerView, Navbar: Navbar},
    },
    {
        path: '/rank', name: 'rank',
        components: {content: RankView, Navbar: Navbar},
    }
];

declare var VueRouter;
declare var Vue;
var router = new VueRouter({
    routes // short for routes: routes
});

router.afterEach((to, from) => {
    var toPath = to.path;
    router.app.active = toPath.split("/")[1];
});

var app = new Vue({
    router
}).$mount('#app');
