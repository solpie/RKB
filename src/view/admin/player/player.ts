import {Profile} from "./profile";
import {CommandId} from "../../Command";
// import {CommandId} from "../../../../event/Command";
export var PlayerView = {
    template: require('./player.html'),
    components: {Profile},
    props: {
        total: {},
        ftId: {},
        ftOptionArr: {
            type: Array,
            required: true,
            default: [{text: 'ft1', value: 1}]
        }
    },
    mounted: function () {
        console.log('player Ready!!');
        this.$http.post('/db/player', {all: true}).then((res)=> {
            console.log(JSON.stringify(res));
            // var a:Array<any> = [];
            var pageCount = 16;
            var count = 0;
            this.countPage = [1];
            this.playerMap = res.data.PlayerMap;
            for (var playerId in res.data.PlayerMap) {
                count++;
                if (count === pageCount) {
                    this.countPage.push(this.countPage.length + 1)
                }
                this.playerArr.push(res.data.PlayerMap[playerId]);
            }
        });

        this.$http.get('/db/ft', (res)=> {
            var ftArr = res.ftArr;
            this.ftOptionArr = [];
            for (var i = 0; i < ftArr.length; i++) {
                var ft = ftArr[i];
                this.ftOptionArr.push({text: ft.name, value: ft.id});
            }
            console.log('ft res:', res, this.ftOptionArr);
        });
    },
    methods: {
        onClearActPlayerGameRec()  {
            this.post(`/panel/stage1v1/${CommandId.cs_clearActPlayerGameRec}`)
        },


        // onSaveToTotalScore() {
        //     this.post(`/panel/stage1v1/${CommandId.cs_saveToTotalScore}`)
        // },
        //
        // onPickPlayer(playerId) {
        //     this.pickPlayerIdArr.push(playerId);
        //     if (this.pickPlayerIdArr.length == 4) {
        //         console.log('pick team');
        //         this.pickPlayerIdArrArr.push(this.pickPlayerIdArr);
        //         this.pickPlayerIdArr = [];
        //     }
        //
        //     this.total = this.pickPlayerIdArrArr.length * 4 + this.pickPlayerIdArr.length;
        // },
        //
        // showFile(files) {
        //
        // },
        //
        // onSubmit(msg) {
        //     console.log('onSubmit', msg)
        // },
        //
        // onAddPlayer() {
        //     ($('#modal-player') as any).openModal();
        //     this.message = "添加球员";
        //     this.isOpen = true;
        //     this.$broadcast(ViewEvent.PLAYER_ADD, {ftOptionArr: this.ftOptionArr.concat()});
        // },
        //
        // onAddPlayerList() {
        //     var a = [];
        //     for (var i = 0; i < this.pickPlayerIdArrArr.length; i++) {
        //         var playerIdArr = this.pickPlayerIdArrArr[i];
        //         a = a.concat(playerIdArr);
        //     }
        //     a = a.concat(this.pickPlayerIdArr);
        //     console.log('playerList', a);
        //     this.post(`/panel/stage1v1/${CommandId.cs_setActPlayer}`, {playerIdArr: a});
        // },
        //
        // onEdit(playerId, event): any {
        //     event.stopPropagation();
        //     console.log("onEdit", playerId);
        //     ($('#modal-player') as any).openModal();
        //     this.message = "编辑球员";
        //     this.$broadcast(ViewEvent.PLAYER_EDIT, {playerId: playerId, ftOptionArr: this.ftOptionArr.concat()});
        // },
        //
        // onFtSelected() {
        //     this.playerArr = [];
        //     for (var playerId in this.playerMap) {
        //         var playerDoc = this.playerMap[playerId];
        //         if (playerDoc.ftId == this.ftId)
        //             this.playerArr.push(playerDoc);
        //     }
        // },

    }

}
