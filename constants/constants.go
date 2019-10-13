package constants

// Requests per second. For value higher than `RPS_VERBOSE` we stop prining individual event.
const RPS int = 1
const RPS_VERBOSE = 1

// Specify event and case study ran
// 1 = Events scenario, case study #1 - raw Data
// 2 = Events scenario, case study #2 - map data + dynamic map for IP addresses
// 3 - Raw text scenario, case study #1 - raw data.
// 4 - Raw text scenario, case study #2 - proto encoded data.
const MODE = 1
