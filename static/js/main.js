
const createHtmlRecord = (jobList) => {
    let messageStream = "";
    console.log(jobList);
    console.log(jobList.length);
    if (jobList.length > 0) {
        jobList[0].messages.forEach((msg) => {
            try {
                const parsedMsg = JSON.parse(msg.body);
                if ("stream" in parsedMsg) {
                    // Build messages
                    messageStream += parsedMsg["stream"];
                } else if ("status" in parsedMsg) {
                    // Push messages
                    const statusText = parsedMsg["status"];

                    if ("id" in parsedMsg) {
                        const imgID = parsedMsg["id"];
                        if ("progress" in parsedMsg) {
                            // Progress bar (in progress)
                            console.log("parsedMsg", parsedMsg); // DEBUG
                            const progressBar = parsedMsg["progress"];
                            messageStream += `${imgID}: ${statusText}\t${progressBar}\n`;
                        } else {
                            // Done pushing
                            messageStream += `${imgID}: ${statusText}\n`;
                        }
                    } else {
                        // Other status
                        messageStream += `${statusText}\n`;
                    }
                } else {
                    console.warn("unable to handle non-stream message", msg);
                }
            } catch (e) {
                // not json
                console.warn("unable to parse JSON for message:", msg);
            }
        });
    }
    console.log("message stream", messageStream);
    return new AnsiUp().ansi_to_html(messageStream);
};

const imgIDToColorPair = imgID => {
    return {
        background: "",
        foreground: ""
    };
};

function invertColor(hex) {
    if (hex.indexOf('#') === 0) {
        hex = hex.slice(1);
    }
    // convert 3-digit hex to 6-digits.
    if (hex.length === 3) {
        hex = hex[0] + hex[0] + hex[1] + hex[1] + hex[2] + hex[2];
    }
    if (hex.length !== 6) {
        throw new Error('Invalid HEX color.');
    }
    // invert color components
    var r = (255 - parseInt(hex.slice(0, 2), 16)).toString(16),
        g = (255 - parseInt(hex.slice(2, 4), 16)).toString(16),
        b = (255 - parseInt(hex.slice(4, 6), 16)).toString(16);
    // pad each with zeros and return
    return '#' + padZero(r) + padZero(g) + padZero(b);
}

function padZero(str, len) {
    len = len || 2;
    var zeros = new Array(len).join('0');
    return (zeros + str).slice(-len);
}


const vm = new Vue({
    el: '#root',
    data() {
        return {
            buildOutput: `<h2>&lt;Build output&gt;</h2>`,
            jobs: []
        }
    },
    created() {
        let ctx = this;
        fetch("/api/jobs")
            .then(function(res) {
                res.json().then(function(resJSON) {
                    if (res.status === 200) {

                        // Attempt to interpret job data
                        ctx.buildOutput = createHtmlRecord(resJSON);

                        // Update jobs list
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
