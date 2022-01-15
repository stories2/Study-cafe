<template>
  <q-page class="q-pa-md">
    <div v-show="pageInit">
      <div class="row">
        <q-list bordered separator>
          <q-item clickable v-ripple v-for="(value, k) in data" v-bind:key="k">
            <q-item-section>
              <q-item-label>Item with caption {{value}}</q-item-label>
              <q-item-label caption>Caption</q-item-label>
            </q-item-section>
            <q-item-section side>
              <q-btn flat icon="cancel" color="red" @click="deleteBtn(k)">
              </q-btn>
            </q-item-section>
          </q-item>
        </q-list>
      </div>
      <div class="row">
        <div class="col">
          <q-btn outline icon="cancel" color="orange" label="hell" class="q-mr-md"></q-btn>
          <q-btn outline icon="cancel" color="gray" label="hell"></q-btn>
          world
        </div>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
// import { Todo, Meta } from 'components/models';
// import ExampleComponent from 'components/CompositionComponent.vue';
import { defineComponent } from 'vue';
import { QBtn } from 'quasar';
import {
      useHellWorld} from './index.code';

export default defineComponent({
  name: 'PageIndex',
  components: { QBtn },
  setup() {
    return { ...useHellWorld() };
  }
});
</script>
