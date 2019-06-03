<template>
    <v-layout column justify-center fill-height>
        <v-flex id="loader" v-if="loading">
            <img src="./../assets/loader.gif">
        </v-flex>
        <v-flex id="errors" v-if="!loading">
            <v-alert :value="response.statusCode === 400" color="error" icon="warning" outline>
                Please enter a search query.
            </v-alert>
            <v-alert :value="response.statusCode === 401" color="error" icon="warning" outline>
                You're not authorized to search.
            </v-alert>
            <v-alert :value="response.statusCode === 500" color="error" icon="warning" outline>
                An unexpected error occurred. Please try again later.
            </v-alert>
        </v-flex>
        <v-flex id="frequencies" v-if="!loading && searchResult.tweets.length > 0">
            <v-layout column>
                <v-flex>
                    <h2 class="headline">Most Frequent Words</h2>
                </v-flex>
                <v-flex>
                    <v-layout row wrap>
                        <v-flex v-for="(frequency, index) in searchResult.frequencies" :key="index">
                            <v-layout column>
                                <v-flex>
                                    <h3>{{ index+1 }}</h3>
                                </v-flex>
                                <v-flex>
                                    <h3>{{ frequency.word }}</h3>
                                </v-flex>
                                <v-flex>
                                    <div>{{ frequency.count }} time(s)</div>
                                </v-flex>
                            </v-layout>
                        </v-flex>
                    </v-layout>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex id="tweets" v-if="!loading && searchResult.tweets.length > 0">
            <v-layout column>
                <v-flex>
                    <h2 class="headline">Tweets</h2>
                </v-flex>
                <v-flex>
                    <v-pagination v-model="page" :length="numPages"></v-pagination>
                </v-flex>
                <v-flex>
                    <v-layout row wrap>
                        <v-flex v-for="(tweet, index) in currentTweets" :key="index">
                            <a class="tweet-link" :href="tweet.url" target="_blank">
                                <v-card class="tweet">
                                    <div>{{ tweet.text }}</div>
                                </v-card>
                            </a>
                        </v-flex>
                    </v-layout>
                </v-flex>
            </v-layout>
        </v-flex>
    </v-layout>
</template>

<script lang="ts">
import {Component, Prop, Watch, Vue} from 'vue-property-decorator'
import {SearchResult, Tweet} from '../assets/js/api'

const resultsPerPage = 50;

@Component
export default class Results extends Vue {
    // INPUT PROPERTIES
    @Prop({required: true})
    private response!: Object;

    @Prop({required: true})
    private searchResult!: SearchResult;

    @Prop({required: true})
    private loading!: boolean;

    // DATA PROPERTIES
    private page: number = 1;

    // FUNCTIONS
    private get numPages(): number {
        // Get the number of pages, given the number of tweets and how many should appear per page.
        return Math.ceil(this.searchResult.tweets.length / resultsPerPage);
    }

    private get currentTweets(): Array<Tweet> {
        // Get the offset of the current page of tweets.
        let offset = resultsPerPage * (this.page-1);

        // Slice the list so that it only contains the current page.
        return this.searchResult.tweets.slice(offset, offset+resultsPerPage);
    }

    // WATCHERS
    @Watch('searchResult.tweets')
    private onTweetsChanged(value: Array<Tweet>, oldValue: Array<Tweet>): void {
        this.page = 1;
    }
}
</script>

<style scoped>
#loader {
    margin: 20px 0px;
    text-align: center;
}
hr {
    margin: 10px 0px;
}
#errors {
    margin: 10px 0px;
}
#frequencies {
    margin: 20px 0px;
}
#tweets {
    margin: 20px 0px;
}
.tweet {
    width: 400px;
    padding: 20px;
    margin: 10px;
}
.tweet-link {
    text-decoration: none;
}
</style>