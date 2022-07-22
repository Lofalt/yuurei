<template>
    <div class="panel">
        <!--        <div class="background" :style="{backgroundImage:background}">-->
        <div class="background" :style="{ backgroundColor: bgc }">
            <!--            <img src="~assets/img/backgroundimg/background1.png"/>-->
        </div>
        <div class="panelTitle">
            <!-- <a @click="archive(article.articleId)">{{ article.articleTitle }}</a> -->hello
        </div>
        <div class="panelSummary">
            <!-- {{ article.articleSummary }} -->这里是简介
        </div>
        <div class="panelInfo">
            <!-- <i class="el-icon-time" /> <span>{{ new Date(article.articleCreateTime).toLocaleDateString() }}</span> -->
            <!--            <i class="el-icon-view"/> <span>阅读 {{article.articleReadTimes}}</span>-->
            <!--            <span>❤ 点赞 {{article.articleLoveTimes}}次</span>-->
            <!-- <i class="el-icon-document" /><span>分类: {{ articleCategories }}</span> -->
        </div>
    </div>
</template>

<script>
// import scrollToTop from "../../util/scrollToTop";
export default {
    name: "Panel",
    //     props: ['article', 'background', 'backgroundColor'],
    //     computed: {
    //         articleCategories() {
    //             let categories = ''
    //             for (let category of this.article.articleCategories) {
    //                 categories += category.categoryName + ' '
    //             }
    //             return categories
    //         },
    //         articleTags() {
    //             let tags = ''
    //             for (let tag of this.article.articleTags) {
    //                 tags += tag.tagName + '  '
    //             }
    //             return tags
    //         },
    //         bgc() {
    //             return this.article.articleCategories[0].categoryColor
    //         }
    //     },
    //     methods: {
    //         archive(articleId) {
    //             scrollToTop(200)
    //             this.$router.push('/archive/' + articleId)
    //             this.$router.push({
    //               path:'/archive',
    //               query:{
    //                 articleId:articleId
    //               }
    //             })
    //         }
    //     }
    // }
}
</script>

<style lang="less" scoped>
.panel {
    background-color: #ffbca1;
    color: white;
    position: relative;
    /*width: 70%;*/
    /*width: 90%;*/
    margin-left: 10px;
    margin-right: 10px;
    margin-bottom: 50px;
    height: 250px;
    border-radius: 7px;
    box-shadow: 0 0 5px rgba(100, 100, 100, .5);
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    background-image: url(../assets/91375949_p0.jpg);
    background-size: cover;

    .panelJump {

        a:visited,
        a:link {
            color: white;
            text-decoration: none;
        }
    }

    .panelSummary {
        position: relative;
        text-overflow: ellipsis;
        text-align: center;
        overflow: hidden;
        white-space: nowrap;
        margin: 20px 30px;
        padding-bottom: 10px;
        /*border-bottom:1px white solid;*/
        transform: translate3d(0, -10px, 0);
        opacity: 0;
        transition: .3s ease all;
        -webkit-transition: .3s ease all;
    }

    &:hover {
        box-shadow: 0 0 10px rgba(252, 168, 167, 0.8);

        .panelTitle {
            transform: translate3d(0, -10px, 0);
            -webkit-transition: .2s ease all;
            font-size: 30px;
            white-space: nowrap;

        }
    }

    &:hover .background {
        transform: scale(1.2, 1.2);
    }

    &:hover .panelSummary {
        transform: translate3d(0, 0, 0);
        -webkit-transition: .3s ease all;
        opacity: 1;
    }

    &:hover .panelInfo {
        transform: translate3d(0, 0, 0);
        opacity: 1;
    }

    .panelInfo {
        position: absolute;
        left: 5px;
        padding-top: 10px;
        padding-right: 100px;
        bottom: 5px;
        margin-left: 2px;
        transform: translate3d(0, -10px, 0);
        opacity: 0;
        transition: .3s ease all;
        -webkit-transition: .3s ease all;

        /*border-top: 1px solid white;*/
        span {
            margin-right: 20px;
        }
    }

    .panelJump {
        position: relative;
    }

    .panelTitle {
        text-align: center;
        position: relative;
        margin-top: -160px;
        font-size: 25px;
        transition: .3s ease all;
        -webkit-transition: .3s ease all;
        cursor: pointer;

        /*span{*/
        /*    cursor: pointer;*/
        /*}*/
        a {
            display: block;
            text-overflow: ellipsis;
            white-space: nowrap;
            overflow: hidden;
        }
    }
}

.background {
    position: relative;
    /*background-image: url(~assets/img/backgroundimg/background1.png);*/
    background-size: cover;
    /*<!--background-position:  0 -200px;-->*/
    width: 100%;
    height: 250px;
    min-height: 250px;
    transition: .3s all ease;
}

@media screen and (max-width: 800px) {
    .panel {
        /*width:100%;*/
        height: 200px;
        font-size: 14px;

        .panelTitle {
            margin-top: -200px;
            opacity: 1;
            transform: translate3d(0, 0, 0);
        }

        .panelInfo {
            opacity: 1;
            transform: translate3d(0, 0, 0);
        }

        .panelSummary {
            opacity: 1;
            transform: translate3d(0, 0, 0);
        }
    }
}
</style>