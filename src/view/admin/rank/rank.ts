import {PlayerInfo} from "../../model/PlayerInfo";
import {mapToArr, descendingProp} from "../../utils/JsFunc";
import {$} from "../../libs";
export var RankView = {
    template: require('./rank.html'),
    props: {
        playerDocArr: {
            type: Array
        }
    },

    mounted() {
        console.log("rank");
        $.post('/db/elo', {}, (data)=> {
            var playerMap = data.PlayerMap;
            this.playerDocArr = mapToArr(playerMap).sort(descendingProp('eloScore'));
            // this.playerDocArr = rank;
        });

    },
    methods: {
        onSortWinPercent() {
            console.log('onSortWinPercent');
        },

        onSortGameCount() {
            for (var p of this.playerDocArr) {
                p.gameCount = PlayerInfo.gameCount(p);
            }
            this.playerDocArr.sort(descendingProp('gameCount'));
            console.log('onSortGameCount');
        }
    }
}
