const vm = new Vue({
    el: '#root',
    data() {
        return {
            jobs: []
        }
    },
    created() {
        let ctx = this;
        fetch("/api/jobs")
            .then(function(res) {
                res.json().then(function(resJSON) {
                    if (res.status === 200) {
                        ctx.jobs = resJSON;
                    }
                }).catch(function() {
                    // Show error
                })
            })
    },
});

var ws = new WebSocket(`wss://${window.location.host}/ws`);
ws.onopen = function (event) {

};

ws.onmessage = function (event) {
    var msg = JSON.parse(event.data);
    console.log(msg);
};
