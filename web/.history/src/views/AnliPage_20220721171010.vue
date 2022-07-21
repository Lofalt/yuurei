  <template>
      <div id="main">
          <img class="test" v-for="(item, index) in picList" :src="item" :key="item"
              :style="{ animationDelay: (index % 4) * 0.1 + 's' }" />
      </div>
  </template>