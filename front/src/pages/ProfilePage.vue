
<template>
    <q-page class="row items-center justify-evenly">
        <q-card class="my-card">
            <q-carousel animated v-model="slide" arrows navigation infinite >
                <q-carousel-slide :name="1" img-src="https://i.pinimg.com/originals/61/2f/d9/612fd974285f8057189b37657542e1ff.jpg" />
                <q-carousel-slide :name="2" img-src="https://cdn.quasar.dev/img/parallax1.jpg" />
                <q-carousel-slide :name="3" img-src="https://cdn.quasar.dev/img/parallax2.jpg" />
                <q-carousel-slide :name="4" img-src="https://cdn.quasar.dev/img/quasar.jpg" />
            </q-carousel>
            <q-card-section>
                <div class="row no-wrap items-center">
                    <q-badge color="green" floating transparent>Online</q-badge>
                    <div class="col text-h6 ellipsis">
                        Jane Doe
                    </div>
                    <div class="col-auto text-grey text-caption q-pt-md row no-wrap items-center">
                        <q-icon name="place" />
                        250 Km
                    </div>
                </div>
                <div class="row no-wrap items-center">
                    <q-rating v-model="ratingModel" size="2em" color="red-7" icon="favorite_border" />
                    <span class="text-caption text-grey q-ml-sm">3.2 (551)</span>
                </div>
                <q-item>
                    <div class="text-caption text-grey">
                        Small description about me <q-badge outline color="blue" label="#fun" /> <q-badge outline color="secondary" label="#code" /> <q-badge outline color="orange" label="#bubbler" />
                    </div>
                </q-item>
                <q-item>
                    <q-btn push color="primary" label="Like" style="margin:10px;" />
                    <q-btn push color="red" label="Block" style="margin:10px;" />
                    <q-btn push color="orange" label="Report" style="margin:10px;" />
                </q-item>
            </q-card-section>
        </q-card>
    </q-page>
</template>

<script>
import { ref } from 'vue'

export default {
  setup () {
    return {
      ratingModel: ref(3),
      slide: ref(1)
    }
  }
}
</script>