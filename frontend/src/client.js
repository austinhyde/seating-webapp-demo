export default {
  async getLocations() {
    return this.apiCall('get_locations');
  },

  async apiCall(call, body) {
    const resp = await fetch('/api/'+call, {
      body: JSON.stringify(body),
    });
    const respBody = await resp.json();
    if (respBody.success) {
      return respBody.data;
    } else {
      throw new Error(respBody.error);
    }
  },
};