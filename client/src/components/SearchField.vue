<template>
    <v-layout row>
        <v-flex>
            <v-layout column>
                <v-flex>
                    <img id="logo" src="./../assets/logo.png">
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex>
            <v-layout column>
                <v-flex>
                    <h1>Tweet Sneak</h1>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex>
            <v-layout column>
                <v-flex>
                    <v-text-field v-model="keywords" solo @keyup.enter="submit"></v-text-field>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex>
            <v-layout column>
                <v-flex>
                    <v-btn @click="submit">Search</v-btn>
                </v-flex>
            </v-layout>
        </v-flex>
    </v-layout>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'vue-property-decorator'

@Component
export default class SearchField extends Vue {
    // INPUT PROPERTIES
    @Prop({required: true})
    private searchAction!: Function;

    // DATA PROPERTIES
    private keywords: string = '';

    // FUNCTIONS
    private submit(): void {
        // Perform the search action with the given keywords.
        this.searchAction(this.keywords);
    }
}
</script>

<style scoped>
    #logo {
        width: 100px;
        height: 75px;
        transition: transform 1s;
        -webkit-transition: transform 1s;
    }
    #logo:hover{
        transform: rotateY(180deg);
        -webkit-transform: rotateY(180deg);
    }
</style>