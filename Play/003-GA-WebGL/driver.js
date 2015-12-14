var container, stats, visuallog;
var camera, controls, scene, renderer;
var objects = [];
var highlightBox;

var mouse = new THREE.Vector2();
var offset = new THREE.Vector3( 10, 10, 10 );

var clock = new THREE.Clock();
var lastIter = 0;
var intervalTime = .1; // seconds
init();
animate();

function init() {

  container = document.getElementById( "container" );
  visuallog = document.getElementById( "visuallog" );

  camera = new THREE.PerspectiveCamera( 70, window.innerWidth / window.innerHeight, 1, 10000 );
  camera.position.z = 1000;

  controls = new THREE.TrackballControls( camera );
  controls.rotateSpeed = 1.0;
  controls.zoomSpeed = 1.2;
  controls.panSpeed = 0.8;
  controls.noZoom = false;
  controls.noPan = false;
  controls.staticMoving = true;
  controls.dynamicDampingFactor = 0.3;

  scene = new THREE.Scene();

  scene.add( new THREE.AmbientLight( 0x555555 ) );
  var light = new THREE.SpotLight( 0xffffff, 1.5 );
  light.position.set( 0, 500, 2000 );
  scene.add( light );

  var geometry = new THREE.Geometry()
  var defaultMaterial = new THREE.MeshLambertMaterial({ color: 0xffffff, shading: THREE.FlatShading, vertexColors: THREE.VertexColors } );

  /*
  var geom = new THREE.BoxGeometry( 1, 1, 1 );
  var color = new THREE.Color();

  var matrix = new THREE.Matrix4();
  var quaternion = new THREE.Quaternion();

  for ( var i = 0; i < 5000; i ++ ) {

    var position = new THREE.Vector3();
    position.x = Math.random() * 10000 - 5000;
    position.y = Math.random() * 6000 - 3000;
    position.z = Math.random() * 8000 - 4000;

    var rotation = new THREE.Euler();
    rotation.x = Math.random() * 2 * Math.PI;
    rotation.y = Math.random() * 2 * Math.PI;
    rotation.z = Math.random() * 2 * Math.PI;

    var scale = new THREE.Vector3();
    scale.x = Math.random() * 200 + 100;
    scale.y = Math.random() * 200 + 100;
    scale.z = Math.random() * 200 + 100;

    quaternion.setFromEuler( rotation, false );
    matrix.compose( position, quaternion, scale );

    // give the geom's vertices a random color, to be displayed

    applyVertexColors( geom, color.setHex( Math.random() * 0xffffff ) );

    geometry.merge( geom, matrix );

    // give the geom's vertices a color corresponding to the "id"

    applyVertexColors( geom, color.setHex( i ) );
  }
  var drawnObject = new THREE.Mesh( geometry, defaultMaterial );
  scene.add( drawnObject );
  */

  renderer = new THREE.WebGLRenderer( { antialias: true } );
  renderer.setClearColor( 0xffffff );
  renderer.setPixelRatio( window.devicePixelRatio );
  renderer.setSize( window.innerWidth, window.innerHeight );
  renderer.sortObjects = false;
  container.appendChild( renderer.domElement );

  stats = new Stats();
  stats.domElement.style.position = 'absolute';
  stats.domElement.style.top = '0px';
  container.appendChild( stats.domElement );

  renderer.domElement.addEventListener( 'mousemove', onMouseMove );

  // for hooking up with genetic.go
  document.driver = new Driver(scene)

}

//

function Driver (scene) {
  this.scene = scene

  // initial empty geometry
  this.geometry = new THREE.Geometry()
  this.defaultMaterial = new THREE.MeshLambertMaterial({ color: 0xffffff, shading: THREE.FlatShading, vertexColors: THREE.VertexColors } )

  this.log = []
}
Driver.prototype.connect = function(iter_fn) {
  console.log("Driver Connected.")
  this.iter_fn = iter_fn
}
Driver.prototype.iter = function() {
  if (typeof(this.iter_fn) === "function") this.iter_fn()
  else console.log("Iter not set up yet")
}
Driver.prototype.addEntity =
function(position_x, position_y, position_z,
          rotation_x, rotation_y, rotation_z,
          scale_x, scale_y, scale_z,
          given_color) {
  var geom = new THREE.BoxGeometry( 1, 1, 1 );
  var color = new THREE.Color();

  var matrix = new THREE.Matrix4();
  var quaternion = new THREE.Quaternion();

  // range -500, 500
  var sideLength = 1000

  var position = new THREE.Vector3();
  position.x = position_x * sideLength - sideLength * .5;
  position.y = position_y * sideLength - sideLength * .5;
  position.z = position_z * sideLength - sideLength * .5;

  /* default until further development */
  var ZERO = 0
  var rotation = new THREE.Euler();
  rotation.x = rotation_x * 2 * Math.PI;
  rotation.y = rotation_y * 2 * Math.PI;
  rotation.z = rotation_z * 2 * Math.PI;

  var scale = new THREE.Vector3();
  scale.x = scale_x * 200 + 100;
  scale.y = scale_y * 200 + 100;
  scale.z = scale_z * 200 + 100;

  quaternion.setFromEuler( rotation, false );
  matrix.compose( position, quaternion, scale );
  // give the geom's vertices a random color, to be displayed

  applyVertexColors( geom, color.setHex( given_color * 0xffffff ) );

  this.geometry.merge( geom, matrix );
}
Driver.prototype.update = function () {
  var drawnObject = new THREE.Mesh( this.geometry, this.defaultMaterial );
  this.scene.add(drawnObject)
  this.scene.remove(this.previousGeneration)
  this.previousGeneration = drawnObject
  this.geometry = new THREE.Geometry()
}

Driver.prototype.visualLog = function() {
  this.log.push(Array.prototype.slice.call(arguments).join(" "))
};
Driver.prototype.visualLogHighlight = function() {
  this.log.push("~" + Array.prototype.slice.call(arguments).join(" ") + "~")
};
Driver.prototype.updateVisualLog = function() {
  visuallog.innerHTML = ""
  var el
  var highlightRE = /^\~(.*)\~$/
  var cont = document.createElement("pre")
  this.log.forEach(function (str) {
    el = document.createElement("div")
    var match = highlightRE.exec(str)
    if (match != null) {
      el.innerText = match[1]
      el.style.backgroundColor = "yellow"
    } else {
      el.innerText = str
    }
    cont.appendChild(el)
  })
  visuallog.appendChild(cont)
  this.log = []
};

function applyVertexColors( g, c ) {

  g.faces.forEach( function( f ) {

    var n = ( f instanceof THREE.Face3 ) ? 3 : 4;

    for( var j = 0; j < n; j ++ ) {

      f.vertexColors[ j ] = c;

    }

  } );

}

function onMouseMove( e ) {

  mouse.x = e.clientX;
  mouse.y = e.clientY;

}

function animate() {

  requestAnimationFrame( animate );

  render();
  stats.update();

}
function render() {
  var driver = document.driver;

  lastIter += clock.getDelta();

  controls.update();

  if (lastIter >= intervalTime) {
    lastIter %= intervalTime;
    lastIter -= intervalTime;
    driver.iter();
  }
  
  renderer.render( scene, camera );
}