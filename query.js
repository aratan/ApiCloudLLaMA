
const API_URL = 'http://192.168.1.22:8080';
const TOKEN_URL = `${API_URL}/token`;
const JOB_URL = `${API_URL}/job`;

async function getToken() {
    const response = await fetch(TOKEN_URL);
    const data = await response.json();
    return data.token;
}

async function createJob(token, phrase) {
    const response = await fetch(JOB_URL, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ phrase })
    });
    const data = await response.json();
    return data.job_id;
}

async function getJobResult(token, jobId) {
    const response = await fetch(`${JOB_URL}?job_id=${jobId}`, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    });
    const data = await response.json();
    return data;
}

async function run() {
    try {
        const token = await getToken();
        const jobId = await createJob(token, 'Hello, world!');
        const jobResult = await getJobResult(token, jobId);
        console.log(jobResult);
    } catch (error) {
        console.error(error);
    }
}

run();
