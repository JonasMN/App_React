const apiKey = 'eiK7P2UC6q3C2rPc4bCksj2mt2IKQ5Uk'


export default function GetGifs({keyword = 'messi'} = {}) {
    const apiURL = `https://api.giphy.com/v1/gifs/search?api_key=${apiKey}&q=${keyword}&limit=10&offset=0&rating=g&lang=en`

    return  fetch(apiURL)
    .then(res => res.json())
    .then(response =>{
      const {data = []} = response
      if (Array.isArray(data)) {
      const gifs = data.map(image => {
          const {images,title,id} = image
          const {url} = image.images.downsized_medium   
            return {title,id,url}
        })
      return gifs
      }
    })
}