from magma.db_service.tests.alembic_testcase import AlembicTestCase

NEW_TABLES = [
    'cbsd_states',
    'domain_proxy_logs',
    'grant_states',
    'request_states',
    'request_types',
    'cbsds',
    'active_mode_configs',
    'channels',
    'requests',
    'grants',
    'responses',
]


class Testb0cad5321c88TestCase(AlembicTestCase):

    def setUp(self) -> None:
        super().setUp()
        self.up_revision = "b0cad5321c88"

    def test_b0cad5321c88_upgrade(self):
        # given / when
        self.upgrade()

        # then
        self.then_tables_are(NEW_TABLES)

    def test_b0cad5321c88_downgrade(self):
        # given
        self.upgrade()

        # when
        self.downgrade()

        # then
        self.then_tables_are()

    def test_downgrade_to_base(self):
        self.verify_downgrade_to_base()
