const inspector = require('inspector');
const fs = require('fs');

const session = new inspector.Session();
session.connect();

session.post('Profiler.enable', () => {
  session.post('Profiler.start', () => {
    // Run the compiler in a timeout so it's on the event loop
    setTimeout(() => {
      try {
        require("../../bin/gopurs.js");
      } catch (e) {
        console.error(e);
      }
    }, 100);

    // Stop after 5 seconds and write profile
    setTimeout(() => {
      session.post('Profiler.stop', (err, { profile }) => {
        if (!err) {
          fs.writeFileSync('./profile.cpuprofile', JSON.stringify(profile));
          console.log("Profile written.");
        }
        process.exit(0);
      });
    }, 5000);
  });
});
