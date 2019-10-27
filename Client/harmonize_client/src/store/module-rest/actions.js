import axios from 'axios'
import config from '../../../config';

export function fetchChannels ({ commit }) {
    axios.get(config.API_URL + 'channels', { crossdomain: true })
        .then(r => r.data)
        .then(response => {
            commit('SET_CHANNELS', response.channels);
        })
}

export function fetchNowPlaying ({ commit }, payload) {
    console.log('received now playing request for ' + payload.channelID);
    axios.get(config.API_URL + 'nowplaying',
        { crossdomain: true, params: { channel: payload.channelID } })
        .then(r => r.data)
        .then(response => {
            console.log(response);
            commit('SET_NOW_PLAYING', response);
        })
}
 