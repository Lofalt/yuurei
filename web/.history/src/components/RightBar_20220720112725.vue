<template>

    <div>这里是右边
        <div class="showpage">
            <router-view></router-view>
        </div>
        <div class="showpage">
            <router-view></router-view>
        </div>
        <div class="showpage">
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
export default {
    name: "RightBar"
}
</script>