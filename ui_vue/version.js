import { promises as fs } from 'fs';

const writeTimestamp = async () => {
    const timestamp = new Date().getTime().toString();
    await fs.writeFile('./public/version.txt', timestamp);
};

writeTimestamp().catch(console.error);
