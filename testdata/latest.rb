cask 'latest' do
  version :latest
  sha256 '5e1e2bcac305958b27077ca136f35f0abae7cf38c9af678f7d220ed0cb51d4f8'

  url "https://example.com/app_#{version}.dmg"
  name 'Example'
  name 'Example (latest)'
  homepage 'https://example.com/'
  license :commercial

  app 'Example (latest).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-latest", target: 'example'
end
