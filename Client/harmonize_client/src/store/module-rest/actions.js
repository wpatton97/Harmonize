import axios from 'axios'

export function fetchChannels ({ commit }) {
    axios.get('http://localhost:3000/api/channels', { crossdomain: true })
        .then(r => r.data)
        .then(response => {
            commit('SET_CHANNELS', response.channels);
        })
}
