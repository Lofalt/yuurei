<template>
    <div class="home">
        <div class="time">
            <div>
                距离那刻
            </div>
            <div>
                大概已经过去了
            </div>
            <div>
                {{ year + '年零' }}
            </div>
            <div>

            </div>
            <div>
                {{ month + '个月' }}
            </div>
            <div>
                {{ +day + '天' }}
            </div>
            <div>
                {{ +hour + '时' }}
            </div>
            <div>
                {{ +minute + '分' }}
            </div>
            <div>
                {{ +second + '秒' }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Home",
    data() {
        return {
            year: '',
            month: '',
            day: '',
            hour: '',
            minute: '',
            second: ''
        }
    },
    mounted() {
        let time = this
        let now = new Date()
        let previous = new Date(2017, 9, 14, 19)
        time.year = Math.floor((now - previous) / (365 * 24 * 60 * 60 * 1000))
        time.month = Math.floor((now - previous) % (365 * 24 * 60 * 60 * 1000) / (30 * 24 * 60 * 60 * 1000))
        previous = (now - previous) % (365 * 24 * 60 * 60 * 1000) % (30 * 24 * 60 * 60 * 1000)
        time.day = Math.floor(previous / (24 * 60 * 60 * 1000))
        previous = previous % (24 * 60 * 60 * 1000)
        time.hour = Math.floor(previous / (60 * 60 * 1000))
        previous = previous % (60 * 60 * 1000)
        time.minute = Math.floor(previous / (60 * 1000))
        previous = previous % (60 * 1000)
        time.second = Math.floor(previous / 1000)
        setInterval(() => {
            let now = new Date()
            let previous = new Date(2017, 9, 14, 19)
            time.year = Math.floor((now - previous) / (365 * 24 * 60 * 60 * 1000))
            time.month = Math.floor((now - previous) % (365 * 24 * 60 * 60 * 1000) / (30 * 24 * 60 * 60 * 1000))
            previous = (now - previous) % (365 * 24 * 60 * 60 * 1000) % (30 * 24 * 60 * 60 * 1000)
            time.day = Math.floor(previous / (24 * 60 * 60 * 1000))
            previous = previous % (24 * 60 * 60 * 1000)
            time.hour = Math.floor(previous / (60 * 60 * 1000))
            previous = previous % (60 * 60 * 1000)
            time.minute = Math.floor(previous / (60 * 1000))
            previous = previous % (60 * 1000)
            time.second = Math.floor(previous / 1000)
        }, 500)
    }
}
</script>

<style lang="less" scoped>
.home {
    position: absolute;
    box-sizing: border-box;
    // min-height: 100vh;
    width: 60%;
    // padding-top: 20%;
    top: 20vh;
    text-align: center;
    background-color: rgb(255, 255, 255);
    font-size: 30px;
    white-space: nowrap;
}

@media screen and (max-width: 800px) {
    // .home {
    //     // padding-top: 45%;
    // }
}

.time {
    // background-color: #ffbca1;
    margin: 0 25%;
    padding: 50px;
    // box-shadow: 0 0 5px rgba(100, 100, 100, .5);
    border-radius: 10px;
    color: black;

    div {
        /*font-family: 幼圆;*/
    }
}

@media screen and (max-width: 800px) {
    .time {
        margin: 0 10%;
    }
}
</style>