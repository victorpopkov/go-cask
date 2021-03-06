cask 'if-global-sha256-last' do
  if MacOS.version <= :leopard
    version '1.0.0'
    url "https://example.com/app_#{version}_mac32.dmg"
  else
    version '2.0.0'
    url "https://example.com/app_#{version}_mac64.dmg"
  end

  sha256 'cd9d7b8c5d48e2d7f0673e0aa13e82e198f66e958d173d679e38a94abb1b2435'
  appcast "https://example.com/sparkle/#{version.major}/appcast.xml"
  name 'Example'
  name 'Example (if-global-sha256-last)'
  homepage 'https://example.com/'

  auto_updates true

  app 'Example (if-global-sha256-last).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
