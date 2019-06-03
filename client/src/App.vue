<template>
  <v-app>
    <v-container>
      <v-layout column>
        <v-flex>
          <Search-Field :searchAction="search"></Search-Field>
        </v-flex>
        <v-flex>
          <hr>
        </v-flex>
        <v-flex>
          <Results :response="response" :searchResult="searchResult" :loading="loading"></Results>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'vue-property-decorator'
import SearchField from './components/SearchField.vue'
import Results from './components/Results.vue'
import http from 'http';
import {DefaultApi, SearchResult} from './assets/js/api'

@Component({
  components: {
    SearchField,
    Results
  }
})
export default class App extends Vue {
  // DATA PROPERTIES
  private response: Object = {statusCode: 200};
  private searchResult: SearchResult = {frequencies: [], tweets: []};
  private loading: boolean = false;
  
  // FUNCTIONS
  private search(keywords: string): void {
    // Create an instance of the client API.
    let api = new DefaultApi(process.env.VUE_APP_API_HOST);

    // Set loading to true.
    this.loading = true;

    // Make a getSearch call to get the tweets for the given keywords.
    api.getSearch(keywords).then((result: {response: http.IncomingMessage, body: SearchResult}) => {
      this.response = {statusCode: result.response.statusCode};
      this.searchResult = result.body;
    }).catch((result: {response: http.IncomingMessage, body: SearchResult}) => {
      this.response = {statusCode: result.response.statusCode};
      this.searchResult = {
        frequencies: [],
        tweets: []
      }
    }).finally(() => {
      this.loading = false;
    });
  }
}
</script>

<style>
</style>
