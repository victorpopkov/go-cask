cask 'if-global-version-last' do
  if MacOS.version <= :leopard
    sha256 'cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435'
    url "https://example.com/app_#{version}_mac32.dmg"
  else
    sha256 '9065ae8493fa73bfdf5d29ffcd0012cd343475cf3d550ae526407b9910eb35b7'
    url "https://example.com/app_#{version}_mac64.dmg"
  end

  version '2.0.0'
  appcast "https://example.com/sparkle/#{version.major}/appcast.xml"
  name 'Example'
  name 'Example (if-global-version-last)'
  homepage 'https://example.com/'

  auto_updates true

  app 'Example (if-global-version-last).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
