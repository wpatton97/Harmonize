import axios from 'axios'
import config from '../../../config'

export function fetchChannels ({ commit }) {
    axios.get(config.API_URL + 'channels', { crossdomain: true })
        .then(r => r.data)
        .then(response => {
            commit('SET_CHANNELS', response);
        })
}

export function fetchNowPlaying ({ commit }, payload) {
    axios.get(config.API_URL + 'nowplaying',
        { crossdomain: true, params: { channel: payload.channelID } })
        .then(r => r.data)
        .then(response => {
            console.log(response);
            commit('SET_NOW_PLAYING', response);
        })
}

export function fetchCandidates ({ commit }, payload) {
    axios.get(config.API_URL + 'candidates',
        { crossdomain: true, params: { channel: payload.channelID } })
        .then(r => r.data)
        .then(response => {
            commit('SET_CANDIDATES', response);
        })
}
 
export function updateActiveChannel ({ commit }, payload) {
    axios.post(config.API_URL + 'channel', null, { crossdomain: true, params: {
        join: payload.channelID
      }})
        .then(r => r.data)
        .then(response => {
        })
}

export function setPlaying({commit}, payload) {
    commit('SET_PLAYING', payload.playing);
}