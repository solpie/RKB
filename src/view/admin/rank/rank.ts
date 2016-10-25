import {PlayerInfo} from "../../model/PlayerInfo";
import {mapToArr, descendingProp} from "../../utils/JsFunc";
import {$} from "../../libs";
import {getEloRank} from "./elo";
export var RankView = {
    template: require('./rank.html'),
    props: {
        playerDocArr: {
            type: Array
        }
    },

    mounted() {
        console.log("rank");
        var gameIdArr = [23, 21, 22, 39];
        var gameDataArr = [];
        var gameId;
        var getGameData = (i)=> {
            if (i < gameIdArr.length) {
                gameId = gameIdArr[i];
                $.get('/api/passerbyking/game/match/' + gameId, (data)=> {
                    // $.get('/db/elo', {gameIdArr: [23, 21, 22, 29, 39]}, (data)=> {
                    console.log(data);
                    gameDataArr.push(data);
                    // this.playerDocArr = rank;
                    getGameData(i + 1);
                });

            }
            else {
                var playerMap = getEloRank(gameDataArr);
                console.log('player map',playerMap);
                this.playerDocArr = mapToArr(playerMap).sort(descendingProp('eloScore'));
            }
        };
        getGameData(0)
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
