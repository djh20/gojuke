import { baseUrl, jsonHeaders } from "./config";

export interface Track {
  id: string;
  name: string;
  fileName: string;
}

export async function fetchTracks(): Promise<Track[]> {
  const tracks: Track[] = await fetch(baseUrl + "/api/tracks").then(r => r.json());
  console.log(tracks);

  return tracks;
}

export async function playTrack(track: Track): Promise<void> {
  console.log(`Attempting to play track: ${track.name}`);

  const reqData = {trackId: track.id};

  await fetch(baseUrl + "/api/playback/queue", {
    method: "POST",
    body: JSON.stringify(reqData),
    headers: jsonHeaders
  });
}

export async function skipTrack(): Promise<void> {
  console.log(`Attempting to skip current track`);
  await fetch(baseUrl + "/api/playback/skip", { method: "POST" });
}