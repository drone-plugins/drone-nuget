const Drone = require('drone-node');
const plugin = new Drone.Plugin();

const Nuget = require('nuget-runner');

const path = require('path');
const shelljs = require('shelljs');

const do_push = function (nuget, workspace, vargs, file) {
  var relative_path = path.relative('.', file);

  console.log('Start pushing ' + file);
  nuget.push(relative_path, {
    source: vargs.source
  }).then(function (stdout) {

   console.log('Successfully pushed ' + file);
  }).catch(function (err) {
    console.error('An error happened while pushing: ' + err);
    process.exit(1);
  });
}

const do_pack_then_push = function (nuget, workspace, vargs, file) {
  console.log('Start packing ' + file);
  nuget.pack({
    spec: file,
    outputDirectory: path.dirname(file)
  }).then(function(stdout) {
    console.log('Successfully packed ' + file);
    var package_path = file.replace('.nuspec', '*.nupkg')
    var resolved_package_file = shelljs.ls(package_path)

    do_push(nuget, workspace, vargs, resolved_package_file[0]);        

  }).catch(function (err) {
    console.error('An error happened while packing: ' + err);
    process.exit(1);
  });
}

const do_upload = function (workspace, vargs) {
  if (vargs.source) {

    var nugetPath = '/usr/lib/nuget/NuGet.exe';
    var nuget = new Nuget({
      apiKey: vargs.api_key,
      verbosity: vargs.verbosity,
      nugetPath: nugetPath
    });

    var nugetVersion = shelljs.exec('mono ' + nugetPath, {silent:true}).head({'-n': 1});
    console.log(nugetVersion);

    var resolved_files = [].concat.apply([], vargs.files.map((f) => { return shelljs.ls(workspace.path + '/' + f); }));
    resolved_files.forEach((file) => {

      if (path.extname(file) === '.nuspec') {
        do_pack_then_push(nuget, workspace, vargs, file);
      } else {
        do_push(nuget, workspace, vargs, file);        
      }
    });
  } else {
    console.error("Parameter missing: NuGet source URL");
    process.exit(1);
  }
}

plugin.parse().then((params) => {

  // gets build and repository information for
  // the current running build
  const build = params.build;
  const repo  = params.repo;
  const workspace = params.workspace;

  // gets plugin-specific parameters defined in
  // the .drone.yml file
  const vargs = params.vargs;

  vargs.verbosity     || (vargs.verbosity = 'quiet');
  vargs.files         || (vargs.files = []);

  do_upload(workspace, vargs);
});
