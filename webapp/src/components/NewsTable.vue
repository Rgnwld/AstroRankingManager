<script setup>
import axios from "axios"

const {data} = await axios.get("http://localhost:8080/public/news")
const newsArr = data
console.log(data)
</script>

<template>
    <div>
        <ol>
            <li v-for="news in newsArr" type="1"><a :href="news.url" target="_blank">{{ news.title}}</a></li>
        </ol>
    </div>
</template>


<style scoped>
</style>