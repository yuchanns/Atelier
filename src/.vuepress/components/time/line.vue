<template>
  <timeline timeline-theme="#31AFB4">
    <template v-for="(title, key) in timeTitles">
      <timeline-title
        :key="key"
        bg-color="#E50743"
      >
        {{ title }}
      </timeline-title>
      <timeline-item
        :bg-color="colors[tkey%6]"
        v-for="(item, tkey) in timeItems[title]"
        :key="title.toString() + tkey"
      >
        <router-link
          :to="item.path"
          :style="{'color': colors[5-tkey%6]}"
        >
          {{ item.title }}
        </router-link>
      </timeline-item>
    </template>
  </timeline>
</template>

<script>
import { Timeline, TimelineItem, TimelineTitle } from 'vue-cute-timeline'
import compareDesc from 'date-fns/compare_desc'
import getYear from 'date-fns/get_year'

export default {
  name: 'TimeLine',

  components: {
    Timeline,
    TimelineItem,
    TimelineTitle,
  },

  data () {
    return {
      num: 0,
      colors: [
        '#F9870F',
        '#E8ED30',
        '#3FA62E',
        '#3BB4D7',
        '#2F4D9E',
        '#71378A',
      ],
    }
  },

  computed: {
    posts () {
      const posts = this.$posts
      return posts.sort((p1, p2) => {
        return compareDesc(p1.createdAt, p2.createdAt)
      })
    },
    year () {
      return {
        end: getYear(this.posts[0].createdAt),
        start: getYear(this.posts[this.posts.length - 1].createdAt),
      }
    },
    timeTitles () {
      return [...Array(this.year.end + 1 - this.year.start).keys()].map(item => this.year.start + item).reverse()
    },

    timeItems () {
      const timelines = {}

      this.timeTitles.forEach(year => {
        timelines[year] = this.posts.filter(item => {
          return (compareDesc(year.toString(), item.createdAt) > -1) && (compareDesc((year + 1).toString(), item.createdAt) < 0)
        })
      })

      return timelines
    },
  },
}
</script>

<style lang="stylus" scoped>
.content
  ul
    padding-left 0
</style>
