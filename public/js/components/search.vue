<template>
<div>
    <h1>搜尋網址</h1>
    <input name="url" v-model="url">
    <button v-on:click="searchUrl">搜尋</button>
    <p v-if="notFoundFlag">{{ notFoundUrl }} 尚未有 Log 資料</p>
</div>
<div>
    <!-- <h4 v-if="urlIsExist()">關鍵字：{{ keyword }}網址：{{ url }}</h4> -->
    <table>
        <tr>
            <td>名次</td>
            <td>紀錄日期</td>
            <td>搜尋引擎</td>
        </tr>
        <tr v-for="seo in seos">
            <td>{{ seo.rating }}</td>
            <td>{{ seo.recorded_at }}</td>
            <td v-if="seo.search_engine=='google'" style="color:blue;">{{ seo.search_engine }}</td>
            <td v-else style="color:purple;">{{ seo.search_engine }}</td>
        </tr>
    </table>
</div>
</template>

<script>
export default {
    data () {
        return {
            keyword: "",
            url: "",
            seos: [],
            notFoundFlag: false,
            notFoundUrl: "",
        }
    },

    methods: {
        searchUrl () {
            if (this.urlIsExist(this.url)) {
                this.renderSeo()
            }
        },

        urlIsExist (url) {
            var data = {url: url}
            this.$http.post('/api/search/url', data, function(res) {
                if (res.length == 0) {
                    this.notFoundUrl = url
                    this.notFoundFlag = true
                } else {
                    this.notFoundFlag = false
                    this.seos = res
                }
            })
            return true
        },

        renderSeo () {
            console.log(this.seos)
        }
    },

    destroyed () {
        clearInterval(this.handle)
    }
}
</script>
