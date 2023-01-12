
<template>
    <div class="q-pa-md row items-start q-gutter-md">
        <div style="max-width: 300px; min-width: 120px; padding-left: 20px;">
            <q-select v-model="SortBy" :options="SortByOptions" label="Sort by" />
        </div>
        <div style="max-width: 300px; min-width: 120px; padding-left: 20px;">
            <q-select v-model="model" :options="options" label="Search by" />
        </div>
        <q-input dark dense standout v-model="text" input-class="text-right" class="q-ml-md">
            <template v-slot:append>
                <q-icon v-if="text === ''" name="search" />
                <q-icon v-else name="clear" class="cursor-pointer" @click="text = ''" />
            </template>
        </q-input>
        <q-slider
        class="q-mt-xl"
        v-model="firstModel"
        color="deep-orange"
        markers
        :marker-labels="kmLabel"
        :min="0"
        :max="10"
        style="max-width:200px;"
        >
        <template v-slot:marker-label-group="scope">
            <div
            v-for="marker in scope.markerList"
            :key="marker.index"
            :class="[ `text-deep-orange-${2 + Math.ceil(marker.value / 2) }`, marker.classes ]"
            :style="marker.style"
            @click="model = marker.value"
            >{{ marker.value }}</div>
        </template>
        </q-slider>

        <div class="q-pa-md row inline items-start q-gutter-md">

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">John Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">Jane Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">Jane Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">Jane Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">Jane Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

            <q-card class="my-card">
                <q-card-section>
                    <div class="text-h6 q-mb-xs">Jane Doe</div>
                </q-card-section>
                <img src="https://cdn.quasar.dev/img/mountains.jpg" style="width: 300px;">
                <q-card-actions align="right">
                    <q-btn flat round color="red" icon="favorite" />
                </q-card-actions>
            </q-card>

        </div>
    </diV>
</template>

<script>
import { ref } from 'vue'

export default {
  setup () {
    return {
      kmLabel: val => `${10 * val}Km`,
      text: ref(''),
      model: ref('Tags'),
      options: [
        'Tags', 'Age', 'Rating', 'Location'
      ],
      SortBy: ref('Tags'),
      SortByOptions: [
        'Tags', 'Age', 'Rating', 'Location'
      ]
    }
  }
}
</script>